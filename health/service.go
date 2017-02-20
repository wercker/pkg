package health

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// New create a new HealthService
func New() *Service {
	return &Service{
		probes: make(map[string]Probe),
	}
}

// The Service periodiablly checks all probes, if any fails it will keep
// track of this. It exposes this information through two handlers,
// livenesProbe and readinessProbe.
type Service struct {
	mu     sync.Mutex
	probes map[string]Probe

	healthyCount   int
	unhealthyCount int
}

// ListenAndServe starts looping through the probes and it will start the
// http server on addr.
func (s *Service) ListenAndServe(addr string) error {
	go s.healthcheck()

	m := http.NewServeMux()

	m.HandleFunc("/live", s.livenessProbe)
	m.HandleFunc("/ready", s.readinessProbe)

	return http.ListenAndServe(addr, m)
}

// healthcheck starts an infinite loop which will iterate over all probes to
// see if there are unhealthy probes.
func (s *Service) healthcheck() {
	// TODO(bvdberg): Support stopping this infinite loop
	for {
		healthy := true

		// iterate all probes
		s.mu.Lock()
		for name, probe := range s.probes {
			// TODO(bvdberg): Add
			err := probe.Healthy()
			if err != nil {
				healthy = false
				log.Printf("Probe %s failed its healthcheck: %+v", name, err)
			}
		}

		sleepDuration := time.Second * 5

		if healthy {
			s.healthyCount++
			s.unhealthyCount = 0
		} else {
			s.unhealthyCount++
			s.healthyCount = 0
			sleepDuration = time.Second * 2
		}

		s.mu.Unlock()

		time.Sleep(sleepDuration)
	}
}

// livenessProbe reports bad health when a probe failed 5 times. It should be
// used to terminate this service.
func (s *Service) livenessProbe(res http.ResponseWriter, req *http.Request) {
	if s.unhealthyCount > 5 {
		http.Error(res, "500 bad health", http.StatusInternalServerError)
		return
	}
}

// readynessProbe reports bad health when a probe failed 1 time. It should be
// used to temporary prevent traffic from coming to this service.
func (s *Service) readinessProbe(res http.ResponseWriter, req *http.Request) {
	if s.unhealthyCount > 0 {
		http.Error(res, "500 bad health", http.StatusInternalServerError)
		return
	}
}

// RegisterProbe adds a new probe to be monitored.
func (s *Service) RegisterProbe(name string, p Probe) {
	s.mu.Lock()

	s.probes[name] = p

	s.mu.Unlock()
}
