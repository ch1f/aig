package aig

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func (s *Suite) TestTransaction() {
	ctx := context.Background()

	s.Run("clean test db", func() {
		err := s.repo.db.Collection(s.repo.collectionName).Drop(ctx)
		s.NoError(err)
	})

	s.Run("ainext: first generate", func() {
		value := s.repo.Next("ainext")
		s.Equal(value, uint64(1))
	})

	s.Run("ainext: second generate", func() {
		value := s.repo.Next("ainext")
		s.Equal(value, uint64(2))
	})

	s.Run("generate: first generate", func() {
		value := s.repo.Generate(ctx, "generate")
		s.Equal(value, uint64(1))
	})

	s.Run("generate: second generate", func() {
		value := s.repo.Generate(ctx, "generate")
		s.Equal(value, uint64(2))
	})

	s.Run("ainext accoring to rules generate", func() {
		value := s.repo.Next("generate")
		s.Equal(value, uint64(3))
	})

	s.Run("generate accoring to rules ainext", func() {
		value := s.repo.Generate(ctx, "ainext")
		s.Equal(value, uint64(3))
	})

	s.Run("consistency ainext", func() {
		tc := 1000

		const X = 100
		var i = 0

		go func(start time.Time) {
			for {
				<-time.After(time.Second)
				ms := time.Now().Sub(start)

				q := float64(i*X) / float64(ms.Seconds())

				fmt.Println("comb/sec:", q, "duration:", ms.Seconds(), "cycles: ", i, "operations", i*X)
			}
		}(time.Now())

		for ; i < tc; i++ {
			g := errgroup.Group{}
			for j := 0; j < X; j++ {

				g.Go(func() error {
					s.repo.Next("consistency_ainext")

					return nil
				})
			}

			_ = g.Wait()
		}

		value := s.repo.Next("consistency_ainext")

		s.Equal(uint64(100001), value) //true result after 1 00 001 times
	})

	// s.Run("consistency generate", func() {
	// 	tc := 1000

	// 	const X = 100
	// 	var i = 0

	// 	go func(start time.Time) {
	// 		for {
	// 			<-time.After(time.Second)
	// 			ms := time.Now().Sub(start)

	// 			q := float64(i*X) / float64(ms.Seconds())

	// 			fmt.Println("comb/sec:", q, "duration:", ms.Seconds(), "cycles: ", i, "operations", i*X)
	// 		}
	// 	}(time.Now())

	// 	for ; i < tc; i++ {
	// 		g := errgroup.Group{}
	// 		for j := 0; j < X; j++ {

	// 			g.Go(func() error {
	// 				s.repo.Generate(ctx, "consistency_generate")

	// 				return nil
	// 			})
	// 		}

	// 		_ = g.Wait()
	// 	}

	// 	value := s.repo.Generate(ctx, "consistency_generate")

	// 	s.Equal(uint64(100001), value) //true result after 1 00 001 times
	// })

}
