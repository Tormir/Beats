package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/km/countbeat/config"
)

// Countbeat configuration.
type SysUsageBeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

type SysUsageData struct {
	User string `json:"USER"`
	Pid string `json:"PID"`
	Cpu string `json:"CPU"`
	Mem string `json:"MEM"`
	Vsz string `json:"VSZ"`
	Rss string `json:"RSS"`
	Tty string `json:"TTY"`
	Stat string `json:"STAT"`
	Start string `json:"START"`
	Time string `json:"TIME"`
	Command string `json:"COMMAND"`
}

// Beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &SysUsageBeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

func (bt *SysUsageBeat) readSysUsageData(path string) (SysUsageData, error) {
	sysusageData := SysUsageData{}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
	return productData, err
	}

	command := exec.Command("python3", dir + "/sysusage.py", "-C", path)
	err = command.Run()
	if err != nil {
		return sysusageData, err
	}

	raw, err := ioutil.ReadFile(dir + "/data.json")
	if err != nil {
		return sysusageData, err 
	}

	err = json.Unmarshal(raw, &sysusageData)
	if err != nil {
		return sysusageData, err

	return sysusageData, nil
}

}

// Run starts countbeat.
func (bt *SysUsageBeat) Run(b *beat.Beat) error {
	logp.Info("SysUsageBeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	counter := 1

	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		sysusageData, err := bt.readSysUsageData(bt.config.Path)

		if err != nil {
			fmt.Println(err)
			return nil
		}

		Pid, _ := strconv.ParseInt(sysusageData.Pid, 64)
		Cpu, _ := strconv.ParseFloat(sysusageData.Cpu, 64)
		Mem, _ := strconv.ParseFloat(sysusageData.Mem, 64)
		Vsz, _ := strconv.ParseInt(sysusageData.Vsz, 64)
		Rss, _ := strconv.ParseInt(sysusageData.Rss, 64)

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":    b.Info.Name,
				"user":    sysusageData.User,
				"pid":     sysusageData.Pid,
				"cpu":     sysusageData.Cpu,
				"mem":     sysusageData.Mem,
				"vsz":     sysusageData.Vsz,
				"rss":     sysusageData.Rss,
				"tty":     sysusageData.Tty,
				"stat":    sysusageData.Stat,
				"start":   sysusageData.Start,
				"time":    sysusageData.Time,
				"command": sysusageData.Command,
			},
		}
		bt.client.Publish(event)
		logp.Info("Event sent")
	}

// Stop stops countbeat.
func (bt *SysUsageBeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
