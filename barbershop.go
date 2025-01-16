1. **Improve readability with named return values and early returns**:
   - Location: `sendBarberHome` function
   - Explanation: Using named return values can make the code more readable by eliminating the need for explicit returns.
   - Before:
     ```go
     func (shop *BarberShop) sendBarberHome(barber string) {
         color.Cyan("%s is going home.", barber)
         shop.BarbersDoneChan <- true
     }
     ```
   - After:
     ```go
     func (shop *BarberShop) sendBarberHome(barber string) {
         done := false
         defer func() {
             if done {
                 shop.BarbersDoneChan <- true
             }
         }()

         color.Cyan("%s is going home.", barber)
         done = true
     }
     ```

2. **Simplify `closeShopForDay` using `sync.WaitGroup`**:
   - Location: `closeShopForDay` function
   - Explanation: Using a `sync.WaitGroup` can simplify the code for waiting for all barbers to finish before closing the shop.
   - Before:
     ```go
     func (shop *BarberShop) closeShopForDay() {
         color.Cyan("Closing shop for the day.")

         close(shop.ClientsChan)
         shop.Open = false

         for a := 1; a <= shop.NumberOfBarbers; a++ {
             <-shop.BarbersDoneChan
         }

         close(shop.BarbersDoneChan)

         color.Green("---------------------------------------------------------------------")
         color.Green("The barbershop is now closed for the day, and everyone has gone home.")
     }
     ```
   - After:
     ```go
     import "sync"

     func (shop *BarberShop) closeShopForDay() {
         color.Cyan("Closing shop for the day.")

         close(shop.ClientsChan)
         shop.Open = false

         var wg sync.WaitGroup
         wg.Add(shop.NumberOfBarbers)

         for i := 0; i < shop.NumberOfBarbers; i++ {
             go func() {
                 defer wg.Done()
                 <-shop.BarbersDoneChan
             }()
         }

         wg.Wait()
         close(shop.BarbersDoneChan)

         color.Green("---------------------------------------------------------------------")
         color.Green("The barbershop is now closed for the day, and everyone has gone home.")
     }
     ```