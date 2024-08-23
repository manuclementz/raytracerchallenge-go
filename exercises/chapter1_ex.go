package exercises

import (
	"fmt"

	rt "raytracerchallenge/core"
)

type Projectile struct {
	Position rt.Tuple
	Velocity rt.Tuple
}

type Environment struct {
	Gravity rt.Tuple
	Wind    rt.Tuple
}

func tick(env Environment, proj Projectile) Projectile {
	position := proj.Position.Add(proj.Velocity)
	velocity := proj.Velocity.Add(env.Gravity).Add(env.Wind)
	return Projectile{position, velocity}
}

func RunSimulation() {
	p := Projectile{rt.NewPoint(0, 1, 0), rt.NewVector(100, 100, 0).Normalize()}
	e := Environment{rt.NewVector(0, -0.01, 0), rt.NewVector(-0.01, 0, 0)}
	fmt.Println("Simulation started")
	step := 1
	for p.Position.Y > 0 {
		p = tick(e, p)
		fmt.Printf("Step %d, Projectile position: %s\n", step, p.Position)
		step++
	}
	fmt.Println("Simulation stopped")
}
