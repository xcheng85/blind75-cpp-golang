package main

// include <iostream>
import (
	"fmt"
)



type GlRenderer struct {
	version string
}

// pointer receiver, golang struct member function
func (c *GlRenderer) Render() error {
	fmt.Printf("GLRenderer: %s", c.version)
	return nil
}

type VulkanRenderer struct {
	version string
}

func (c *VulkanRenderer) Render() error {
	fmt.Printf("VulkanRenderer: %s", c.version)
	return nil
}
