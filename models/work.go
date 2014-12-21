package models

import (
	"html/template"
	"regexp"
	"strings"
)

var slugRexp = map[*regexp.Regexp]string{
	regexp.MustCompile("[^a-z0-9]+"): "-",
	regexp.MustCompile("^-+"):        "",
	regexp.MustCompile("-+$"):        "",
}

type Tech string

func (self Tech) Slug() string {
	slug := strings.ToLower(string(self))
	for re, repl := range slugRexp {
		slug = re.ReplaceAllString(slug, repl)
	}
	return slug
}

type WorkType int

const (
	WorkTypePersonal WorkType = iota
	WorkTypeProfessional
	WorkTypeFreelance
)

type Work struct {
	Image       string
	Type        WorkType
	Name        string
	Description template.HTML
	Techs       []Tech
	Url         string
	Duration    int
}

func AllWorks() []*Work {
	works := make([]*Work, 0)
	works = append(works, &Work{
		"/static/img/neilgarb_co_za.png",
		WorkTypePersonal,
		"neilgarb.co.za",
		template.HTML(`This website, built in barebones Go. With a bit of gorilla. The frontend is built with vanilla jQuery.  <a href="https://github.com/NeilGarb/www">Fork me</a>!`),
		[]Tech{"Go"},
		"http://neilgarb.co.za/",
		1,
	})
	works = append(works, &Work{
		"/static/img/yadda_co_za.png",
		WorkTypePersonal,
		"yadda.co.za",
		template.HTML(`A group buying aggregator,  <em>yadda</em> would poll 20+ South African group buying and daily deal websites and send subscribed users a daily digest email.`),
		[]Tech{"PHP", "Zend Framework", "MySQL"},
		"",
		3,
	})
	works = append(works, &Work{
		"/static/img/myqron_co_za.png",
		WorkTypePersonal,
		"myqron.co.za",
		template.HTML(`Myqron acted as a translator between pull services such as RSS and push services, in this case web hooks.  It would poll the feeds you configured and post to a selected address when a new entry was found.`),
		[]Tech{"PHP", "MySQL"},
		"",
		2,
	})
	works = append(works, &Work{
		"/static/img/americanflat.png",
		WorkTypeFreelance,
		"Americanflat",
		template.HTML(`Built an image processor in Python to compose artworks and various media into web and print-ready images using ImageMagick.  For example, an artwork might be composed with a blank canvas to create an image of the artwork on the canvas.`),
		[]Tech{"Python", "ImageMagick", "RabbitMQ"},
		"http://www.americanflat.com/",
		2,
	})
	works = append(works, &Work{
		"/static/img/udotest.png",
		WorkTypeFreelance,
		"UDoTest",
		template.HTML(`Built simple e-commerce website offering private home HPV tests. The platform integrates with with courier APIs, testing labs and diagnosing doctors.`),
		[]Tech{"Python", "MySQL"},
		"https://www.udotest.com/",
		4,
	})
	works = append(works, &Work{
		"/static/img/kaluma.png",
		WorkTypeProfessional,
		"Kaluma",
		"Senior PHP developer, working on core services. The Kaluma platform is a process-orientated workflow management tool.",
		[]Tech{"PHP", "MariaDB", "RabbitMQ"},
		"http://www.kaluma.com/",
		10,
	})
	works = append(works, &Work{
		"/static/img/superbalist.png",
		WorkTypeProfessional,
		"Superbalist",
		"Rebuilt API using Laravel. The superbalist.com website and backoffice are powered by this.",
		[]Tech{"PHP", "Laravel", "MySQL"},
		"http://superbalist.com/",
		8,
	})
	works = append(works, &Work{
		"/static/img/takealot.png",
		WorkTypeProfessional,
		"TAKEALOT",
		"Introduced Python into the TAKEALOT stack and it has become the language of choice for their core services. Worked on various projects, including the management of data imports and ETL (20m+ records updated on a daily and weekly basis).",
		[]Tech{"Python", "PHP", "MySQL", "MongoDB", "RabbitMQ", "Solr", "AWS"},
		"http://www.takealot.com/",
		13,
	})
	works = append(works, &Work{
		"/static/img/zoopy.png",
		WorkTypeProfessional,
		"Zoopy",
		"Lead developer at Zoopy for three years. The platform changed shape a number of times, and the stack changed accordingly.  Built the API, transcoding platform and backoffice all in PHP/ZF. The platform was hosted in AWS and we used Rightscale to manage our deployment.",
		[]Tech{"PHP", "Zend Framework", "MySQL", "Video transcoding", "AWS", "Solr"},
		"",
		42,
	})
	works = append(works, &Work{
		"/static/img/oi.png",
		WorkTypeProfessional,
		"Ogilvy Interactive",
		"As a senior PHP developer I worked primarily on the VW account. Dabbled in digital strategy.",
		[]Tech{"PHP", "CakePHP", "MySQL"},
		"http://oi.co.za/",
		12,
	})
	works = append(works, &Work{
		"/static/img/go2africa.png",
		WorkTypeProfessional,
		"Go2Africa",
		"Built G2A's PHP platform, TravelQuest, using CakePHP. The platform had to cater for a strong in-house SEO and content team, and integrate with the business's booking software.",
		[]Tech{"PHP", "CakePHP", "MySQL"},
		"http://www.go2africa.com/",
		13,
	})
	works = append(works, &Work{
		"/static/img/iol.png",
		WorkTypeProfessional,
		"Independent Online",
		"As a senior developer, I worked on a number of projects including IOLJobs and the INMSA (IOL's parent company) image library, used as an archive for photos after they were published (photos were indexed using their embedded IPTC profiles).",
		[]Tech{"PHP", "MySQL"},
		"http://www.iol.co.za/",
		30,
	})
	works = append(works, &Work{
		"/static/img/flickswitch.png",
		WorkTypeFreelance,
		"Flickswitch",
		"Built a basic vanilla PHP frontend for Flickswitch.",
		[]Tech{"PHP"},
		"http://www.flickswitch.co.za/",
		1,
	})
	works = append(works, &Work{
		"/static/img/bitcoin.png",
		WorkTypeFreelance,
		"- Under NDA -",
		"Built a full-stack PHP/MySQL application with frontend and admin. Includes a fully-functional Bitcoin wallet, as well as some custom cryptographic work.",
		[]Tech{"PHP", "MariaDB", "Bitcoin", "RabbitMQ"},
		"",
		6,
	})
	works = append(works, &Work{
		"/static/img/avision.png",
		WorkTypeFreelance,
		"AVIsion",
		"Built a hardware inventory management tool for AVIsion.  The back-office tool allowed staff around the country to register new hardware and log/acknowledge faults.",
		[]Tech{"PHP", "MySQL"},
		"",
		2,
	})
	return works
}

// WorkTechs returns a map whose keys are Techs and whose values are relative
// durations for which the technologies have been used, normalized to [0, 100].
func WorkTechs() map[Tech]float64 {
	ret := make(map[Tech]float64, 0)
	works := AllWorks()

	// First iterate over works and determine what the max is.
	tmp := make(map[Tech]int, 0)
	max := 0
	var found bool
	for _, work := range works {
		for _, tech := range work.Techs {
			_, found = tmp[tech]
			if found {
				tmp[tech] += work.Duration
			} else {
				tmp[tech] = work.Duration
			}
			if tmp[tech] > max {
				max = tmp[tech]
			}
		}
	}

	// Now iterate over tmp and normalize.
	for tech, duration := range tmp {
		ret[tech] = float64(duration) / float64(max) * float64(100)
	}

	return ret
}
