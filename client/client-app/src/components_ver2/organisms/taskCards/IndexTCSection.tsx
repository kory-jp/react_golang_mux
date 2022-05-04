import { Box } from "@mui/system";
import { push } from "connected-react-router";
import { FC, useCallback, useEffect, useLayoutEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";
import useLoadingState from "../../../hooks/useLoadingState";
import usePagination from "../../../hooks/usePagination";
import useReturnTop from "../../../hooks/useReturnTop";
import { nowLoadingState } from "../../../reducks/loading/actions";
import { RootState } from "../../../reducks/store/store";
import { indexTaskCards } from "../../../reducks/taskCards/operations";
import { TaskCards } from "../../../reducks/taskCards/types";
import LoadingLayout from "../../molecules/loading/LoadingLayout";
import DefaultPagination from "../../molecules/pagination/DefaultPagination";
import TaskCardItem from "./TaskCardItem";

type Params = {
  id: string | undefined
}

export const IndexTCSection: FC = () => {
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const loadingState = useLoadingState()
  const {sumPage, setSumPage, queryPage} = usePagination()
  const returnTop = useReturnTop()
  
  useLayoutEffect(() => {
    dispatch(nowLoadingState(true))
  }, [])
  
  useEffect(() => {
    dispatch(indexTaskCards(id, setSumPage, queryPage))
  }, [id, setSumPage, queryPage])
  const taskCards: TaskCards = useSelector((state: RootState) => state.taskCards)

  const onChangeCurrentPage = useCallback((event: React.ChangeEvent<unknown>, page: number) => {
    dispatch(push(`/todo/show/${id}?page=${page}`))
    returnTop()
  }, [])

  return(
    <>
      {
        loadingState ? (
          <LoadingLayout />
        ) : (
          <>
            <Box
              className='taskCard__index'
            >
              {
                taskCards !== null && taskCards.length > 0 ? (
                  <Box
                    className='tc__items'
                  >
                    {
                      taskCards.map(taskCard => (
                        <Box
                          key={taskCard.id}
                          sx={{
                            marginBottom: {
                              xs: '24px',
                              md: '40px',
                            }
                          }}
                        >
                          <TaskCardItem 
                            taskCard={taskCard}
                          />                          
                        </Box>
                      ))
                    }
                  </Box>
                ) : (
                  <>
                    <Box
                      color='#FFF'
                      sx={{
                        fontSize: {
                          xs: '16px',
                          md: '24px',
                        }
                      }}
                    >
                      タスクカードの投稿はありません
                    </Box>
                  </>
                )
              }
              <DefaultPagination 
                count={sumPage}
                onChange={onChangeCurrentPage}
                page={queryPage}
              />
            </Box>
          </>
        )
      }
    </>
  )
}

export default IndexTCSection;