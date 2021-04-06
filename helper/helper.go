package helper

import (
	"bytes"
	"io"
	"os"
	"text/template"
)

const JOB_TEMPLATE = `
job "{{.Name}}" {
  datacenters = ["dc1"]

  group "{{.Name}}" {
    network {
      port "containerPort" {
        to = {{.ContainerPort}}
      }
    }

    task "{{.Name}}"
      driver = "docker"

      config {
        image = {{.Image}}
        ports = ["containerPort"]
      }
    }
  }
}
`

type Job struct {
	Name          string
	Image         string
	ContainerPort int
}

func (j Job) doTemplate(out io.Writer) error {
	t, err := template.New("myTemplate").Parse(JOB_TEMPLATE)
	if err != nil {
		return err
	}
	err = t.Execute(out, j)
	if err != nil {
		return err
	}
	return nil
}

func (j Job) RenderJob() (string, error) {
	var tpl bytes.Buffer
	err := j.doTemplate(&tpl)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil

}

func (j Job) BuildJob(filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = j.doTemplate(outFile)
	if err != nil {
		return err
	}
	return nil
}
