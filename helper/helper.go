package helper

import (
	"bytes"
	"io"
	"os"
	"text/template"

	"github.com/angrycub/d2n/petname"
	"github.com/spf13/cobra"
)

const JOB_TEMPLATE = `
{{- $D1 := "  " -}}{{- $D2 := "    " -}}{{- $D3 := "      " -}}
job "{{.Name}}" {
  datacenters = ["dc1"]

  group "{{.Name}}" {
    task "{{.Name}}"
      driver = "docker"

      config {
        image = "{{.Image}}"
{{- if .Command }}
        command = "{{.Command}}"
{{- end -}}
{{- if .Entrypoint }}
        entrypoint = [{{ range $I, $K :=.Entrypoint }}{{ if ne $I 0}},{{ end }}"{{$K}}"{{ end }}]
{{- end -}}
{{- if .Arguments }}
        args = [{{ range $I, $K :=.Arguments }}{{ if ne $I 0}},{{ end }}"{{$K}}"{{ end }}]
{{- end -}}
{{- if .Privileged }}
        privileged = true
{{- end -}}
{{- if .Tty }}
        tty = true
{{- end -}}
{{- if .Interactive }}
        interactive = true
{{- end }}
      }
    }
  }
}
`

func (c Config) doTemplate(out io.Writer) error {
	t, err := template.New("myTemplate").Parse(JOB_TEMPLATE)
	if err != nil {
		return err
	}
	err = t.Execute(out, c)
	if err != nil {
		return err
	}
	return nil
}

func (c Config) RenderJob() (string, error) {
	//	fmt.Println("In RenderJob")
	var tpl bytes.Buffer
	err := c.doTemplate(&tpl)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil

}

func (c Config) BuildJob(filename string) error {
	//	fmt.Println("In BuildJob")
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = c.doTemplate(outFile)
	if err != nil {
		return err
	}
	return nil
}

func AddSharedDockerFlags(cmd *cobra.Command, config *Config) {

	// --publish , -p
	// -p 127.0.0.1:80:8080/tcp
	//cmd.Flags().StringArrayVarP(&config.exposePorts, "expose", "p", []string{}, "Expose a port or a range of ports")

	// --env , -e
	// -e SERVER_MODE=production
	cmd.Flags().StringArrayVarP(&config.Env, "env", "e", []string{}, "Expose a port or a range of ports")

	// --entrypoint
	// --entrypoint "/bin/sh"
	cmd.Flags().StringArrayVar(&config.Entrypoint, "entrypoint", []string{}, "Overwrite the default ENTRYPOINT of the image")

	// --name
	// --name "fooName"
	cmd.Flags().StringVar(&config.Name, "name", "", "Overwrite the default ENTRYPOINT of the image")

	// --privileged
	cmd.Flags().BoolVar(&config.Privileged, "privileged", false, "Give extended privileges to this container")

	// --interactive , -i
	cmd.Flags().BoolVarP(&config.Interactive, "interactive", "i", false, "Keep STDIN open even if not attached")

	// --tty , -t
	cmd.Flags().BoolVarP(&config.Tty, "tty", "t", false, "Allocate a pseudo-TTY")

}

func (jobConfig *Config) ParseArguments(args []string) {
	if jobConfig.Name == "" {
		jobConfig.Name = petname.Generate(3, "-")
	}

	jobConfig.Image = args[0]

	if len(args) >= 2 {
		jobConfig.Command = args[1]
	}

	if len(args) >= 3 {
		jobConfig.Arguments = args[2:]
	}
}
