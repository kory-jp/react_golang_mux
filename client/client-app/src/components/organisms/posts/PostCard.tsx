import { Button, Card, CardActions, CardContent, CardMedia, Typography } from "@mui/material";
import { push } from "connected-react-router";
import { useCallback, VFC } from "react";
import { useDispatch } from "react-redux";

import sample1 from "../../../assets/images/sample1.jpeg"
import { Todo } from "../../../reducks/todos/types";

type Props = {
  todo: Todo,
}

export const PostCard: VFC<Props> = (props) => {
  const {todo} = props;
  const dispatch = useDispatch()
  const imagePath = `http://localhost:8000/api/img/${todo.imagePath}`

  const onclickToShowTodo = useCallback(() => {
    dispatch(push(`/todo/show/${todo.id}`))
  }, [])

  return(
    <>
      <Card>
        <CardMedia 
          component="img"
          // height="180"
          image={todo.imagePath? imagePath : sample1}
          sx={{
            height : {
              sx: "120px",
              md: "200px"
            },
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
          <Button>
            アクション
          </Button>
        </CardActions>
      </Card>
    </>
  )
}

export default PostCard;