import { Box } from "@mui/system";
import { FC, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";
import useLoadingState from "../../../hooks/useLoadingState";
import usePagination from "../../../hooks/usePagination";
import { RootState } from "../../../reducks/store/store";
import { indexTaskCards } from "../../../reducks/taskCards/operations";
import { TaskCards } from "../../../reducks/taskCards/types";

type Params = {
  id: string | undefined
}

export const IndexTCSection: FC = () => {
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const loadingState = useLoadingState()
  const {sumPage, setSumPage, queryPage} = usePagination()

  useEffect(() => {
    dispatch(indexTaskCards(id, setSumPage, queryPage))
  }, [id])
  const taskCards: TaskCards = useSelector((state: RootState) => state.taskCards)

  console.log(taskCards)

  return(
    <>
      <Box
        color='#FFF'
      >
        Hello
      </Box>
    </>
  )
}

export default IndexTCSection;