import { CardMedia, Grid } from "@mui/material";
import { Box } from "@mui/system";
import axios from "axios";
import { push } from "connected-react-router";
import { FC, useCallback, useEffect, useLayoutEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useParams } from "react-router-dom";
import taskComment from "../../../assets/images/taskComment.svg"
import useLoadingState from "../../../hooks/useLoadingState";
import usePagination from "../../../hooks/usePagination";
import useReturnTop from "../../../hooks/useReturnTop";
import { nowLoadingState } from "../../../reducks/loading/actions";
import { RootState } from "../../../reducks/store/store";
import { indexTaskCards } from "../../../reducks/taskCards/operations";
import { TaskCards } from "../../../reducks/taskCards/types";
import { pushToast } from "../../../reducks/toasts/actions";
import { PrimaryButton } from "../../atoms/buttons/PrimaryButton";
import LoadingLayout from "../../molecules/loading/LoadingLayout";
import DefaultPagination from "../../molecules/pagination/DefaultPagination";
import CreateTCModal from "./CreateTCModal";
import TaskCardItem from "./TaskCardItem";

type Params = {
  id: string | undefined
}

type Response = {
  status: number,
  message: string,
  incompleteTaskCount: number,
}

export const IndexTCSection: FC = () => {
  const dispatch = useDispatch()
  const params: Params = useParams();
  const id: number = Number(params.id)
  const loadingState = useLoadingState()
  const {sumPage, setSumPage, queryPage} = usePagination()
  const [openCreateTCModal, setOpenTCModal] = useState(false)
  const [incompleteTaskCardCount, setIncompleteTaskCardCount] = useState(0)
  const returnTop = useReturnTop()

  const getIncompleteTackCardCount = useCallback((id: number) => {
    const apiURL = process.env.REACT_APP_API_URL + `taskcard/incompletetaskcount/${id}`
    axios
      .get(apiURL,
          {
            withCredentials: true,
            headers:{
              'Accept': 'application/json',  
              'Content-Type': 'multipart/form-data'
            }          
          }
        ).then((response) => {
          const resp: Response = response.data
          console.log(response)
          if (resp.status == 200) {
            setIncompleteTaskCardCount(resp.incompleteTaskCount)
          } else {
            dispatch(pushToast({title: response.data.message, severity: "error"}))             
          }
        })
        .catch((error) => {
          dispatch(pushToast({title: 'データ取得に失敗しました', severity: "error"}))
        })
  }, [])
  
  useLayoutEffect(() => {
    dispatch(nowLoadingState(true))
    getIncompleteTackCardCount(id)
  }, [])
  
  useEffect(() => {
    dispatch(indexTaskCards(id, setSumPage, queryPage))
  }, [id, setSumPage, queryPage])
  const taskCards: TaskCards = useSelector((state: RootState) => state.taskCards)

  const onChangeCurrentPage = useCallback((event: React.ChangeEvent<unknown>, page: number) => {
    dispatch(push(`/todo/show/${id}?page=${page}`))
    returnTop()
  }, [])

  const onClickOpenCreateTCModal = useCallback(()=> {
    setOpenTCModal(true)
  }, [])

  const onClickCloseCreateTCModal = useCallback(() => {
    setOpenTCModal(false)
  }, [])

  console.log(incompleteTaskCardCount)

  return(
    <>
      <Box
        className='taskCard__heading'
        sx={{
          backgroundColor: '#2D2A2A',
          borderRadius: "10px",
          padding: '16px',
          marginBottom: {
            xs: '40px',
          }
        }}
      >
        <Box
          className='button'
          sx={{
            marginBottom: {
              xs: '16px',
            }
          }}
        >
          <PrimaryButton
            onClick={onClickOpenCreateTCModal}
          >
            タスクカードを作成
          </PrimaryButton>
        </Box>
        <Grid
          container
          spacing={{xs: '2', md: '0'}}
        >
          <Grid>
            <CardMedia
              component="img"
              image={taskComment}
              sx={{
                height : {
                  xs: 'auto'
                },
                width: {
                  xs: '320px',
                },
              }}
            />
          </Grid>
          <Grid
            sx={{
              marginX: 'auto'
            }}
          >
            <Box
              sx={{
                marginBottom: {
                  xs: '16px'
                }
              }}
            >
              残りのタスクカード
            </Box>
            <Box
              sx={{
                fontSize: {
                  xs: '24px',
                  md: '40px',
                },
                textAlign: 'center'
              }}
            >
              {incompleteTaskCardCount}
            </Box>
          </Grid>
        </Grid>
      </Box>    
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
      <CreateTCModal 
        open={openCreateTCModal}
        onClose={onClickCloseCreateTCModal}
      />
    </>
  )
}

export default IndexTCSection;