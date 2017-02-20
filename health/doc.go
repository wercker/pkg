/*
  Package health provides a simple health monitoring server.

  To create a health service simply create a new HealthService using New,
  register any probes using the RegisterProbe function and finally calling
  ListenAndServe in a separate go routine:

      go func() {
        healthService := health.New()

        probe := createProbe()
        healthService.RegisterProbe("sample", probe)

        // This will block unless something failed
        err := healthService.ListenAndServe(":7000")
      }()

  To implement a Probe simply adhere to the Probe interface. Use type assertion
  when using a different interface:

      store := createStore() // <- store contains a interface which does not implement health.Probe
      if probe, ok := interface{}(store).(health.Probe); ok {
        healthService.RegisterProbe("store", probe)
      }

  Install this package with govendor:

      govendor fetch -v github.com/wercker/health

  Install this package during a wercker run:

      build:
        steps:
          - script:
              name: install govendor
              code: go get -u github.com/kardianos/govendor

          - script:
              name: force "go get" over ssh
              code: git config --global url."git@github.com:".insteadOf "https://github.com/"

          - add-ssh-key:
              keyname: WALTERBOT

          - add-to-known_hosts:
              hostname: github.com
              fingerprint: 16:27:ac:a5:76:28:2d:36:63:1b:56:4d:eb:df:a6:48
              type: rsa

          - script:
              name: install depdendencies
              code: govendor sync

  MongDB session probe, including trying to recover from an io.EOF error:

      type MongoStore struct {
        session *mgo.Session
      }

      func (s *MongoStore) Healthy() error {
        err := s.session.Ping()
        if err != nil {
          if err == io.EOF {
            s.session.Refresh()
          }

          return err
        }

        return nil
      }
*/
package health
