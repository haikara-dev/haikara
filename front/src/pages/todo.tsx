import type { NextPage } from "next";
import Head from "next/head";
import { Container, Typography, Box, Stack, Card } from "@mui/material";
import React, { useEffect, useState } from "react";
import AddTodoForm from "@/components/AddTodoForm";
import TodoRow from "@/components/TodoRow";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import { useAuthUserContext } from "@/lib/AuthUser";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;

export const TodoStatusIncomplete: number = 0;
export const TodoStatusCompleted: number = 1;

export type Todo = {
  id: number;
  text: string;
  status: number;
};

const Todos: NextPage = () => {
  const [data, setData] = useState<Todo[]>([]);
  const [isLoading, setLoading] = useState(false);
  const { authUser } = useAuthUserContext();

  const getRequestHeaders = async () => {
    const idToken = await authUser?.getIdToken();
    return {
      Authorization: `Bearer ${idToken}`,
    };
  };

  const loadData = async () => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(BACKEND_API_URL + "/todo", {
        method: "GET",
        headers: headers,
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      const json = await res.json();
      setData(json);
    } catch (err) {
      console.log(err);
    }
    setLoading(false);
  };

  const addTodo = async (text: string) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(BACKEND_API_URL + "/todo", {
        method: "POST",
        headers: {
          ...headers,
          ...{
            "Content-Type": "application/json",
          },
        },
        body: JSON.stringify({
          text: text,
        }),
      });
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const doneTodo = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_API_URL + "/todo/"),
        {
          method: "PUT",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
          body: JSON.stringify({
            status: TodoStatusCompleted,
          }),
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const undoTodo = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_API_URL + "/todo/"),
        {
          method: "PUT",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
          body: JSON.stringify({
            status: TodoStatusIncomplete,
          }),
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  const removeTodo = async (id: number) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_API_URL + "/todo/"),
        {
          method: "DELETE",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };
  const updateTodo = async (id: number, text: string) => {
    try {
      const headers = await getRequestHeaders();
      const res = await fetch(
        new URL(id.toString(), BACKEND_API_URL + "/todo/"),
        {
          method: "PUT",
          headers: {
            ...headers,
            ...{
              "Content-Type": "application/json",
            },
          },
          body: JSON.stringify({
            text: text,
          }),
        }
      );
      if (!res.ok) throw new Error(`${res.status}: ${res.statusText}`);
      await loadData();
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    setLoading(true);
    loadData();
  }, []);

  return (
    <div>
      <Head>
        <title>TODO</title>
        <meta name="description" content="TODO" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Header />
      <Box
        component="main"
        sx={{
          height: "100vh",
        }}
      >
        <Container
          sx={{
            p: 2,
          }}
        >
          <Typography variant="h3" component="h1">
            TODO
          </Typography>
          <Box>
            <AddTodoForm addTodo={addTodo} />
          </Box>
          {isLoading ? (
            <div>Loading...</div>
          ) : (
            <Stack gap={2} mt={2} pr={8}>
              {data.map((todo) => {
                return (
                  <Card key={todo.id}>
                    <TodoRow
                      key={todo.id}
                      todo={todo}
                      doneTodo={doneTodo}
                      undoTodo={undoTodo}
                      removeTodo={removeTodo}
                      updateTodo={updateTodo}
                    />
                  </Card>
                );
              })}
            </Stack>
          )}
        </Container>
      </Box>
      <Footer />
    </div>
  );
};

export default Todos;
