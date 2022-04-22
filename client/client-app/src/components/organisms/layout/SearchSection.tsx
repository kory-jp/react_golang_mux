import { Grid, Paper } from "@mui/material";
import { Box } from "@mui/system";
import { push } from "connected-react-router";
import { FC, useCallback, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../../reducks/store/store";
import { indexTags } from "../../../reducks/tags/operations";
import { Tag, Tags } from "../../../reducks/tags/types";
import TagSelction from "../../molecules/tag/TagSelction";


export const SearchSection: FC = () => {
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(indexTags())
  }, [])

  const options: Tags = useSelector((state: RootState) => state.tags)

  const onChangeToTagPage = useCallback((event: Tag) => {
    console.log(event)
    const tagId: number = event.id
    dispatch(push(`/todo/tag/${tagId}`))
  },[])

  return (
    <Box
      id="searchSection"
    >
      <Paper
        sx={{
          padding: {
            sm: '10px',
            md: '20px'
          },
          width: '90%',
          marginX: 'auto'
        }}
      >
        <Grid 
          container
          padding='20px'
        >
          <Grid
            item
            minWidth="300px"
          >
            <TagSelction 
              isMulti={false} 
              options={options} 
              onChange={onChangeToTagPage}
            />
          </Grid>
        </Grid>
      </Paper>
    </Box>
  )
}