package config

import "strconv"

type Config struct {
	Data map[string]string
}

func (c *Config) String(k string) string {
	v, _ := c.Data[k]
	return v
}

func (c *Config) Int(k string) int {
	s, _ := c.Data[k]
	v, _ := strconv.Atoi(s)
	return v
}
