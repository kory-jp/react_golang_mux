import { Grid, Paper } from "@mui/material";
import { Box } from "@mui/system";
import { push } from "connected-react-router";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import { RootState } from "../../../reducks/store/store";
import { indexTags } from "../../../reducks/tags/operations";
import { Tag, Tags } from "../../../reducks/tags/types";
import { Todo} from "../../../reducks/todos/types";
import { makeOptions } from "../../../utils/makeOptions";
import { PrimaryButton } from "../../atoms/button/PrimaryButton";
import TagSelction from "../../molecules/tag/TagSelection";


export const SearchSection: FC = () => {
  const dispatch = useDispatch()
  const [tag, setTag] = useState(0)
  const [importance, setImportance] = useState(0)
  const [urgency, setUrgency] = useState(0)

  useEffect(() => {
    dispatch(indexTags())
  }, [])
  const options: Tags = useSelector((state: RootState) => state.tags)
  const { importanceOptions, urgencyOptions } = makeOptions()

  const onChangeSelectTags = useCallback((event: Tag) => {
    setTag(event.id)
  }, [setTag])

  const onChangeImportance = useCallback((event: Todo) => {
    console.log(typeof(event))
    setImportance(event.id)
  }, [setImportance])

  const onChangeUrgency = useCallback((event: Todo) => {
    setUrgency(event.id)
  }, [setUrgency])

  const onClickToSearchPage = useCallback(() => {
    dispatch(push(`/todo/search?tagId=${tag}&importance=${importance}&urgency=${urgency}`))
  },[tag, importance, urgency])

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
              onChange={onChangeSelectTags}
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
              onChange={onChangeImportance}
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
              onChange={onChangeUrgency}
            />
          </Grid>
          <Grid
            item
            minWidth="300px"
            marginRight="20px"
            marginBottom="20px"
          >
            <PrimaryButton
              onClick={onClickToSearchPage}
            >
              検索
            </PrimaryButton>
          </Grid>
        </Grid>
      </Paper>
    </Box>
  )
}