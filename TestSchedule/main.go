package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {
	var (
		cronJob    *CronJob
		now        time.Time
		nextTime   time.Time
		expr       *cronexpr.Expression
		cronTables map[string]*CronJob
	)

	cronTables = make(map[string]*CronJob)

	expr, _ = cronexpr.Parse("*/5 * * * * * *")

	now = time.Now()
	nextTime = expr.Next(now)

	cronJob = &CronJob{
		expr:     expr,
		nextTime: nextTime,
	}

	cronTables["job1"] = cronJob

	expr, _ = cronexpr.Parse("*/5 * * * * * *")

	now = time.Now()
	nextTime = expr.Next(now)

	cronJob = &CronJob{
		expr:     expr,
		nextTime: nextTime,
	}

	cronTables["job2"] = cronJob

	go func() {
		for {
			var (
				cronJob *CronJob
				now     time.Time
				jobName string
			)

			now = time.Now()

			for jobName, cronJob = range cronTables {
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					go func(jobName string) {
						fmt.Println("Do ", jobName)
					}(jobName)

					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println("nextTime:", cronJob.nextTime)
				}
			}

			time.Sleep(100 * time.Millisecond)
		}

	}()

	time.Sleep(100 * time.Second)
}
