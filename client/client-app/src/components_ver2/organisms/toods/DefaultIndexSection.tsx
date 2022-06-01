import { Grid } from "@mui/material";
import { Box } from "@mui/system";
import { FC } from "react";
import { Todos } from "../../../reducks/todos/types";
import LoadingLayout from "../../molecules/loading/LoadingLayout";
import DefaultPagination from "../../molecules/pagination/DefaultPagination";
import TodoCard from "./TodoCard";

type Props = {
  todos: Todos,
  sumPage: number,
  setSumPage: React.Dispatch<React.SetStateAction<number>>,
  queryPage: number,
  loadingState: boolean,
  onChangeCurrentPage: (event: React.ChangeEvent<unknown>, page: number) => void,
}

export const DefaultIndexSection: FC<Props> = (props) => {
  const {todos, sumPage, setSumPage, queryPage, loadingState, onChangeCurrentPage} = props

  return(
    <>
      {
        loadingState ? (
          <LoadingLayout />
        ) : (
          <Box>
            {
              todos != null && todos.length > 0 && (
                <Grid
                  container
                  spacing={4}
                  sx={{
                    justifyContent: {
                      md: 'space-between'
                    }
                  }}
                >
                  {
                    todos.map (todo => (
                      <Grid
                        item
                        key={todo.id}
                        sx={{
                          width: {
                            xs: '100%',
                            md: 'auto',
                          },
                        }}
                      >
                        <TodoCard 
                          todo={todo}
                          setSumPage={setSumPage}
                          queryPage={queryPage}
                        />
                      </Grid>
                    ))
                  }
                </Grid>
              )
            }
          <DefaultPagination 
            count={sumPage}
            onChange={onChangeCurrentPage}
            page={queryPage}
          />
          </Box>
        )
      }
    </>
  )
}

export default DefaultIndexSection;