import React, { useState } from "react";
import { Todo, TodoStatusCompleted, TodoStatusIncomplete } from "@/pages/todo";
import { Box, Checkbox, IconButton, Typography } from "@mui/material";

import DeleteIcon from "@mui/icons-material/Delete";
import EditTodoForm from "@/components/EditTodoForm";
export type TodoRowProps = {
  todo: Todo;
  doneTodo: (id: number) => void;
  undoTodo: (id: number) => void;
  removeTodo: (id: number) => void;
  updateTodo: (id: number, text: string) => void;
};

const TodoRow: React.FC<TodoRowProps> = ({
  todo,
  doneTodo,
  undoTodo,
  removeTodo,
  updateTodo,
}) => {
  const [isEditting, setEditting] = useState<boolean>(false);

  const onChangeCheckboxHandler = (e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    if (todo.status === TodoStatusIncomplete) {
      doneTodo(todo.id);
    } else {
      undoTodo(todo.id);
    }
  };

  const onClickRemoveHandler = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    removeTodo(todo.id);
  };

  const onClickTextHandler = (e: React.MouseEvent<HTMLDivElement>) => {
    e.preventDefault();
    setEditting(true);
  };

  const onEndEditHandler = () => {
    setEditting(false);
  };

  return (
    <Box display="flex" alignItems="center">
      <Checkbox
        checked={todo.status === TodoStatusCompleted}
        onChange={onChangeCheckboxHandler}
      />
      <Box
        onClick={onClickTextHandler}
        sx={{
          flexGrow: 1,
        }}
      >
        {isEditting ? (
          <EditTodoForm
            todo={todo}
            updateTodo={updateTodo}
            onEndEdit={onEndEditHandler}
          />
        ) : todo.status === TodoStatusCompleted ? (
          <Typography variant="body1" sx={{ textDecoration: "line-through" }}>
            {todo.text}
          </Typography>
        ) : (
          <Typography variant="body1">{todo.text}</Typography>
        )}
      </Box>

      <IconButton onClick={onClickRemoveHandler} aria-label="remove">
        <DeleteIcon />
      </IconButton>
    </Box>
  );
};

export default TodoRow;
