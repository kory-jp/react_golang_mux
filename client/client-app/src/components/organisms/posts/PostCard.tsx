import { push } from "connected-react-router";
import { Button, Card, CardActions, CardContent, CardMedia, Checkbox, FormControlLabel, Grid, Link, Typography } from "@mui/material";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import DeleteIcon from '@mui/icons-material/Delete';
import StickyNote2Icon from '@mui/icons-material/StickyNote2';

import sample1 from "../../../assets/images/sample1.jpeg"
import { Todo } from "../../../reducks/todos/types";
import { deleteTodo, updateIsFinished } from "../../../reducks/todos/operations";

type Props = {
  todo: Todo,
}

export const PostCard: FC<Props> = (props) => {
  const {todo} = props;
  const dispatch = useDispatch()
  const [finish, setFinish] = useState(false)
  const imagePath = `http://localhost:8000/api/img/${todo.imagePath}`

  useEffect(() => {
    setFinish(todo.isFinished)
  }, [])

  const onclickToShowTodo = useCallback(() => {
    dispatch(push(`/todo/show/${todo.id}`))
  }, [])

  const onChangeIsFinished = useCallback(() => {
    if (finish) {
      setFinish(false)
      dispatch(updateIsFinished(todo.id, false))
    } else {
      setFinish(true)
      dispatch(updateIsFinished(todo.id, true))
    }
  }, [todo, finish])

  const onClickDelete = useCallback(() => {
    dispatch(deleteTodo(todo.id))
  }, [todo])

  return(
    <>
      <Card
        sx={{
          transition: '0.7s',
          bgcolor: finish? 'text.disabled' : 'white'
        }}
      >
        <CardMedia 
          component="img"
          image={todo.imagePath? imagePath : sample1}
          sx={{
            height : {
              sx: "120px",
              md: "200px"
            },
            transition: '0.7s',
            filter: finish? 'grayscale(100%)' : '',
            '&:hover': {
              cursor: 'pointer'
            }
          }}
          onClick={onclickToShowTodo}
        />
        <CardContent>
          <Typography
            sx={{
              '&:hover': {
                cursor: 'pointer'
              }
            }}
            onClick={onclickToShowTodo}
          >
            {todo.title}
          </Typography>
        </CardContent>
        <CardActions>
          <Grid container>
            <Grid item xs={6}>
              <FormControlLabel
                control={<Checkbox 
                            checked={finish}
                            value={finish}
                            onChange={onChangeIsFinished}
                          />} 
                label="finish"
                sx={{
                  fontSize: '8px',
                  marginBottom: '3px'
                }}
              />
            </Grid>
            <Grid item xs={6}>
              <Button
                onClick={onClickDelete}
                sx={{
                  color: 'black',
                  fontSize: '8px',
                  marginBottom: '3px'
                }}
              >
                <DeleteIcon />
                 Delete
              </Button>
              <Button
                onClick={onclickToShowTodo}
                sx={{
                  color: 'black',
                  fontSize: '8px'
                }}
              >
                <StickyNote2Icon />
                more
              </Button>
            </Grid>
          </Grid>
        </CardActions>
      </Card>
    </>
  )
}

export default PostCard;