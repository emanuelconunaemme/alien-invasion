# SAGA Alien invasion
Golang solution for the SAGA Alien Invasion problem.

## Build and Run
From root directory
- `go build`
- `./aliens --aliensNumber=<N> --inputFile=<input_file_path>` with `N` as the number of aliens playing the game.

Alternative `make run` for faster iterations with the default params.

### Input File
It's a txt file with the following format:
```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
```
Assumption: the data should be consistent (e.g.: if `Foo north=Bar`, there should also be `Bar south=Foo`), otherwise it will throw an exception

## Run Tests
Just run `make test` or `go test <local_test_dir>` to run specific tests.

## Further Optimizations
For sake of simplicity there is not a lot of redundancy. Although, for improved efficiency it's advised to introduce redundant data structures to avoid sequential access of large arrays. E.g.: keep a pointer in the aliens list of the current city for each alien to avoid iterating through the city to find the alien.