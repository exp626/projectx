package generator

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
)

type Command struct {
	CommandCode byte         `json:"command_code"`
	Name        string       `json:"name"`
	Direction   string       `json:"direction"`
	Body        ProtocolType `json:"body"`
}

const commandMethodFmt = `%s(ctx context.Context, body %s) (err error)`

func (c *Command) FormatMethod() (formatted string, err error) {
	formatted = fmt.Sprintf(commandMethodFmt+"\n", c.Name, c.Body.Name)

	return formatted, nil
}

func (c *Command) RenameAsLanguage(lang OutputLanguage) (err error) {
	switch lang {
	case GoLanguage:
		c.Body.Name.ToCamel()
		c.Name = strcase.ToCamel(c.Name)
	}

	switch c.Body.Type {
	case structType:
		opt, ok := c.Body.Options.(*StructOptions)
		if !ok {
			return errors.New("invalid struct options")
		}

		switch lang {
		case GoLanguage:
			for i := 0; i < len(opt.Fields); i++ {
				opt.Fields[i].Name = strcase.ToCamel(opt.Fields[i].Name)
				opt.Fields[i].Type.ToCamel()
			}
		}
	}

	return nil
}

const serverCommandCodeCaseFmt = `case CommandCode%s:
		{
			if len(rawCommandBody) < Size%s {
				return protocol.ErrBodyIsTooShort
			}
	
			body, err := New%s([Size%s]byte(rawCommandBody))
			if err != nil {
				return err
			}
	
			err = s.service.%s(context.Background(), body)
			if err != nil {
				return err
		
			}
		}
`

func (c *Command) FormatCommandCodeCase() (formatted string, err error) {
	formatted = fmt.Sprintf(
		serverCommandCodeCaseFmt,
		c.Name,
		c.Body.Name,
		c.Body.Name,
		c.Body.Name,
		c.Name,
	)

	return formatted, nil
}

const commandCodeFmt = `CommandCode%s byte = %d`

func (c *Command) FormatCommandCode() (formatted string, err error) {
	formatted = fmt.Sprintf(commandCodeFmt+"\n", c.Name, c.CommandCode)

	return formatted, nil
}

const clientCommandFmt = `
func (c *Client)%s(ctx context.Context, body %s) (err error) {
	rawCommandBody, err := New%sBytes(body)
	if err != nil {
		return err
	}

	rawBody := make([]byte, 0, Size%s+1)

	rawBody = append(rawBody, CommandCode%s)
	rawBody = append(rawBody, rawCommandBody[:]...)

	n, err := c.wr.Write(rawBody)
	if err != nil {
		return err
	}

	if n != len(rawBody){
		return protocol.ErrAllInformationWasNotWritten
	}

	return nil
}
`

func (c *Command) FormatClientCommandMethod() (formatted string, err error) {
	formatted = fmt.Sprintf(
		clientCommandFmt,
		c.Name,
		c.Body.Name,
		c.Body.Name,
		c.Body.Name,
		c.Name,
	)

	return formatted, nil
}

// COMMANDS TYPE AND METHODS

type Commands []Command

func (c Commands) FormatCommandsBodyTypes() (formatted string, err error) {
	formatted += "\n"

	for _, item := range c {
		typeDeclaration, err := item.Body.Format()
		if err != nil {
			return formatted, err
		}

		formatted += typeDeclaration
	}

	formatted += "\n"

	return formatted, nil
}

func (c Commands) FormatInterfaceMethods() (formatted string, err error) {
	formatted += "\n"

	for _, item := range c {
		typeDeclaration, err := item.FormatMethod()
		if err != nil {
			return formatted, err
		}

		formatted += typeDeclaration
	}

	formatted += "\n"

	return formatted, nil
}

const commandCodesFmt = `const (%s)`

func (c Commands) FormatCommandCodes() (formatted string, err error) {
	commandCodesDeclaration := "\n"

	for _, item := range c {
		commandCodeDeclaration, err := item.FormatCommandCode()
		if err != nil {
			return formatted, err
		}

		commandCodesDeclaration += commandCodeDeclaration
	}

	commandCodesDeclaration += "\n"

	formatted = fmt.Sprintf(commandCodesFmt, commandCodesDeclaration)

	return formatted, nil
}

func (c Commands) FormatCommandCodeServerCases() (formatted string, err error) {
	formatted = "\n"

	for _, item := range c {
		commandCodeCaseDeclaration, err := item.FormatCommandCodeCase()
		if err != nil {
			return formatted, err
		}

		formatted += commandCodeCaseDeclaration
	}

	return formatted, nil
}

const commandsFileFmt = `
	import (
		"context"
		%s
	)
		
	%s

	type Service interface {
		%s
	}

	%s
`

func (c Commands) Format() (formatted string, err error) {
	commandCodesDeclaration, err := c.FormatCommandCodes()
	if err != nil {
		return formatted, err
	}

	typesDeclaration, err := c.FormatCommandsBodyTypes()
	if err != nil {
		return formatted, err
	}

	serviceInterfaceDeclaration, err := c.FormatInterfaceMethods()
	if err != nil {
		return formatted, err
	}

	formatted = fmt.Sprintf(
		commandsFileFmt,
		func() string {
			if strings.Contains(typesDeclaration, "protocol.") {
				return "protocol \"github.com/exp626/projectx/pkg/protocol\""
			}

			return ""
		}(),
		commandCodesDeclaration,
		serviceInterfaceDeclaration,
		typesDeclaration,
	)

	return formatted, nil
}

const serverFmt = `
import (
	"context"
	"errors"
	protocol "github.com/exp626/projectx/pkg/protocol"
)

type Server struct{	
	service Service
}

func NewServer(service Service) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) HandleCommand(rawBody []byte) (err error){
	if len(rawBody) < 2 {
		return protocol.ErrBodyIsTooShort
	}

	commandCode := rawBody[0]

	rawCommandBody := rawBody[1:]

	switch commandCode{%s
	default:
		return errors.New("unknown command code")
	}

	return nil
}
`

func (c Commands) FormatServer() (formatted string, err error) {
	serverCommandCodeCases, err := c.FormatCommandCodeServerCases()
	if err != nil {
		return formatted, err
	}

	formatted = fmt.Sprintf(serverFmt, serverCommandCodeCases)

	return formatted, nil
}

func (c Commands) FormatClientCommandMethods() (formatted string, err error) {
	formatted += "\n"

	for _, item := range c {
		clientCommandFormatted, err := item.FormatClientCommandMethod()
		if err != nil {
			return formatted, err
		}

		formatted += clientCommandFormatted
	}

	formatted += "\n"

	return formatted, nil
}

const clientFmt = `
	import (
		"io"
		"context"
		protocol "github.com/exp626/projectx/pkg/protocol"
	)
	
	type Client struct {
		wr io.Writer
	}
	
	func NewClient(wr io.Writer) *Client {
		return &Client{
			wr: wr,
		}
	}

	%s
`

func (c Commands) FormatClient() (formatted string, err error) {
	clientMethodsDeclaration, err := c.FormatClientCommandMethods()
	if err != nil {
		return formatted, err
	}

	formatted = fmt.Sprintf(clientFmt, clientMethodsDeclaration)

	return formatted, nil
}
