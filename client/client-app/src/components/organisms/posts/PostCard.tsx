import { Button, Card, CardActions, CardContent, CardMedia, Typography } from "@mui/material";

import sample1 from "../../../assets/images/sample1.jpeg"

export const PostCard = () => {
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
            タイトル
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