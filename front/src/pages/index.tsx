import React, { ReactElement } from "react";
import DefaultLayout from "@/components/layouts/DefaultLayout";
import { NextPageWithLayout } from "@/pages/_app";
import Typography from "@mui/material/Typography";
import { Counter } from "@/features/counter/Counter";
import { useGetPokemonByNameQuery } from "@/services/pokemon";

const Home: NextPageWithLayout = () => {
  // Using a query hook automatically fetches data and returns query values
  const { data, error, isLoading } = useGetPokemonByNameQuery("bulbasaur");
  // Individual hooks are also accessible under the generated endpoints:
  // const { data, error, isLoading } = pokemonApi.endpoints.getPokemonByName.useQuery('bulbasaur')

  return (
    <div>
      <Typography variant="h3" component="h1">
        haikara
      </Typography>
      <Counter></Counter>

      {error ? (
        <>Oh no, there was an error</>
      ) : isLoading ? (
        <>Loading...</>
      ) : data ? (
        <>
          <h3>{data.species.name}</h3>
        </>
      ) : null}
    </div>
  );
};

Home.getLayout = (page: ReactElement) => <DefaultLayout>{page}</DefaultLayout>;

export default Home;
