import { Box } from "@mui/system";
import { useDispatch, useSelector } from "react-redux";
import { RooState } from "../../../reducks/store/store";

import PostCard from "../../organisms/posts/PostCard";

export const IndexTodos = () => {
  const dispatch = useDispatch()
  const user = useSelector((state: RooState) => state.user)
  const todos = useSelector((state: RooState) => state.todos)
  return(
    <>
      <Box
        sx={{
          marginY: {
            xs: '40px',
            md: '80px'
          },
          marginX: {
            xs: '10px',
            md: '50px'
          }
        }}
      >
        <PostCard />
      </Box>
    </>
  )
}

export default IndexTodos;