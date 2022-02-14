import { Button, Card, CardActions, CardContent, CardMedia, Typography } from "@mui/material";
import { VFC } from "react";

import sample1 from "../../../assets/images/sample1.jpeg"
import { Todo } from "../../../reducks/todos/types";

type Props = {
  todo: Todo,
}

export const PostCard: VFC<Props> = (props) => {
  const {todo} = props;
  return(
    <>
      <Card
        sx={{
          maxWidth: '345px'
        }}
      >
        <CardMedia 
          component="img"
          height="140"
          image={sample1}
        />
        <CardContent>
          <Typography>
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