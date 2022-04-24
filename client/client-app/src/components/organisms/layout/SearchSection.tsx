import { Grid, Paper } from "@mui/material";
import { Box } from "@mui/system";
import { push } from "connected-react-router";
import { FC, useCallback, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../../reducks/store/store";
import { indexTags } from "../../../reducks/tags/operations";
import { Tag, Tags } from "../../../reducks/tags/types";
import { makeOptions } from "../../../utils/makeOptions";
import TagSelction from "../../molecules/tag/TagSelection";


export const SearchSection: FC = () => {
  const dispatch = useDispatch()

  useEffect(() => {
    dispatch(indexTags())
  }, [])

  const options: Tags = useSelector((state: RootState) => state.tags)
  const { importanceOptions, urgencyOptions } = makeOptions()
  const onChangeToTagPage = useCallback((event: Tag) => {
    const tagId: number = event.id
    dispatch(push(`/todo/tag/${tagId}`))
  },[])

  return (
    <Box
      id="searchSection"
      marginBottom="20px"
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
            marginRight="20px"
            marginBottom="20px"
          >
            <TagSelction
              placeholder={"タグを選択してください"}
              isMulti={false} 
              options={options} 
              onChange={onChangeToTagPage}
            />
          </Grid>
          <Grid
            item
            minWidth="300px"
            marginRight="20px"
            marginBottom="20px"
          >
            <TagSelction
              placeholder={"重要度を選択してください"}
              isMulti={false} 
              options={importanceOptions} 
              onChange={onChangeToTagPage}
            />
          </Grid>
          <Grid
            item
            minWidth="300px"
            marginRight="20px"
            marginBottom="20px"
          >
            <TagSelction
              placeholder={"緊急度を選択してください"} 
              isMulti={false} 
              options={urgencyOptions} 
              onChange={onChangeToTagPage}
            />
          </Grid>
        </Grid>
      </Paper>
    </Box>
  )
}