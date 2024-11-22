package main

import (
	"fmt"
	"jobs-daemon/internal/types"
	"log"

	"github.com/charmbracelet/huh"
)

var (
	job          types.ArcJob
	jobName      string
	jobCmd       string
	jobConfig    string
	newJobConfig string
	action       string
	updateJob    bool
	confirm      bool
)

func main() {
	for {
		formHome := huh.NewForm(

			huh.NewGroup(
				huh.NewSelect[string]().
					Options(huh.NewOptions("Insert", "List & Update", "Quit")...).
					Value(&action).
					Title("Choose an action").
					Height(5),
			),
		)

		err := formHome.Run()
		if err != nil {
			log.Fatal(err)
		}

		switch action {
		case "Insert":
			form_insert := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Job name").
						Value(&jobName),
					huh.NewSelect[string]().
						Options(huh.NewOptions("finetune-text", "finetune-multimodal")...).
						Value(&jobCmd).
						Title("Job type").
						Height(5),
					huh.NewInput().
						Title("Job config").
						Value(&jobConfig),
					huh.NewConfirm().
						Title("Confirm?").
						Value(&confirm),
				),
			)

			err = form_insert.Run()
			if err != nil {
				log.Fatal(err)
			}

			if confirm {
				newArcJob := types.ArcJob{
					JobId:     -1,
					JobName:   jobName,
					JobCmd:    jobCmd,
					JobConfig: jobConfig,
					Started:   0,
					Done:      0,
					Failed:    0,
				}
				err := InsertArcJob(newArcJob)
				if err != nil {
					log.Fatal(err)
				}
			}

		case "List & Update":
			formJobs := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[types.ArcJob]().
						Title("Queued jobs").
						Value(&job).
						Height(16).
						OptionsFunc(func() []huh.Option[types.ArcJob] {
							jobs, err := ListArcJobs()
							if err != nil {
								log.Fatal(err)
							}
							return huh.NewOptions(jobs...)
						}, &job),
					huh.NewConfirm().
						Title("Update this job config?").
						Value(&updateJob),
				),
			)

			err = formJobs.Run()
			if err != nil {
				log.Fatal(err)
			}

			if updateJob {
				updateMsg := fmt.Sprintf("Type an updated config to replace %s", job.JobConfig)
				formUpdate := huh.NewForm(
					huh.NewGroup(
						huh.NewInput().
							Title(updateMsg).
							Value(&newJobConfig),
						huh.NewConfirm().
							Title("Confirm?").
							Value(&confirm),
					),
				)

				err = formUpdate.Run()
				if err != nil {
					log.Fatal(err)
				}

				if confirm {
					var newArcJob = types.ArcJob{
						JobId:     job.JobId,
						JobName:   job.JobName,
						JobCmd:    job.JobCmd,
						JobConfig: newJobConfig,
						Started:   job.Started,
						Done:      job.Done,
						Failed:    job.Failed,
					}
					fmt.Printf("updating arc job --> %v", newArcJob)
					err = UpdateArcJob(newArcJob)
					if err != nil {
						log.Fatal(err)
					}
				}
			}

		case "Quit":
			return
		}
	}
}
