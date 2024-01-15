package service_test

import (
	"app/app/internal"
	"app/app/internal/repository"
	"app/app/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGet = func() (t map[int]internal.TicketAttributes, err error) {
			t = map[int]internal.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "johndoe@gmail.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTotalAmountTickets()

		// assert
		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}

// Test for ServiceTicketDefault.GetTicketsAmountByDestinationCountry
func TestServiceTicketDefault_GetTicketsAmountByDestinationCountry(t *testing.T) {
	t.Run("success to get tickets amount by destination country", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGetTicketByDestinationCountry = func(country string) (t map[int]internal.TicketAttributes, err error) {
			t = map[int]internal.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "juandoe@example.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
				2: {
					Name:    "Mark",
					Email:   "markdoe@example.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetTicketsAmountByDestinationCountry("USA")

		// assert
		expectedTotal := 2
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})

}

// Test for ServiceTicketDefault.GetPercentageTicketsByDestinationCountry
func TestServiceTicketDefault_GetPercentageTicketsByDestinationCountry(t *testing.T) {
	t.Run("success to get percentage tickets by destination country", func(t *testing.T) {
		// arrange
		// - repository: mock
		rp := repository.NewRepositoryTicketMock()
		// - repository: set-up
		rp.FuncGet = func() (t map[int]internal.TicketAttributes, err error) {
			t = map[int]internal.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "juandoe@example.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
				2: {
					Name:    "Mark",
					Email:   "markdoe@example.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
				3: {
					Name:    "Mark",
					Email:   "markdoe@example.com",
					Country: "Argentina",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}
		rp.FuncGetTicketByDestinationCountry = func(country string) (t map[int]internal.TicketAttributes, err error) {
			t = map[int]internal.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "juandoe@example.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
				2: {
					Name:    "Mark",
					Email:   "markdoe@example.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		// - service
		sv := service.NewServiceTicketDefault(rp)

		// act
		total, err := sv.GetPercentageTicketsByDestinationCountry("USA")

		// assert
		expectedTotal := 100.0 * 2.0 / 3.0
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)

	})
}
