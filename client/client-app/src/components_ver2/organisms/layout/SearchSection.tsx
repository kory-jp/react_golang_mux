import { Grid } from "@mui/material";
import { Box } from "@mui/system";
import { push } from "connected-react-router";
import { FC, useCallback, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import useReturnTop from "../../../hooks/useReturnTop";
import { RootState } from "../../../reducks/store/store";
import { indexTags } from "../../../reducks/tags/operations";
import { Tag, Tags } from "../../../reducks/tags/types";
import { Todo } from "../../../reducks/todos/types";
import { makeOptions } from "../../../utils/makeOptions";
import { PrimaryButton } from "../../atoms/buttons/PrimaryButton";
import TagSelection from "../../molecules/tag/TagSelection";

export const SearchSection: FC = () => {
  const dispatch = useDispatch()
  const [tag, setTag] = useState(0)
  const [importance, setImportance] = useState(0)
  const [urgency, setUrgency] = useState(0) 
  const returnTop = useReturnTop()

  useEffect(() => {
    dispatch(indexTags())
  }, [dispatch])

  const options: Tags = useSelector((state: RootState) => state.tags)
  const { importanceOptions, urgencyOptions } = makeOptions()

  const onChangeSelectTags = useCallback((event: Tag) => {
    setTag(event.id)
  }, [setTag])

  const onChangeImportance = useCallback((event: Todo) => {
    setImportance(event.id)
  }, [setImportance])

  const onChangeUrgency = useCallback((event: Todo) => {
    setUrgency(event.id)
  }, [setUrgency])

  const onClickToSearchPage = useCallback(() => {
    dispatch(push(`/todo/search?tagId=${tag}&importance=${importance}&urgency=${urgency}`))
    returnTop()
  },[dispatch, returnTop, tag, importance, urgency])

  return(
    <>
      <Box
        bgcolor='#2D2A2A'
        sx={{
          paddingY: {
            xs: '24px',
          },
          paddingX: {
            xs: '16px',
          },
        }}
        borderRadius='10px'
      >
        <Grid
          container
          spacing={2}
          sx={{
            justifyContent: {
              xs: 'center',
              sm: 'center',
              md: 'center',
              lg: 'space-between',
            }
          }}
        >
          <Grid item>
            <TagSelection 
              placeholder={'タグを選択してください'}
              isMulti={false}
              options={options}
              onChange={onChangeSelectTags}
            />
          </Grid>
          <Grid item>
            <TagSelection 
              placeholder={'重要度を選択してください'}
              isMulti={false}
              options={importanceOptions}
              onChange={onChangeImportance}
            />
          </Grid>
          <Grid item>
            <TagSelection 
              placeholder={'緊急度を選択してください'}
              isMulti={false}
              options={urgencyOptions}
              onChange={onChangeUrgency}
            />
          </Grid>
          <Grid item>
            <Box
              sx={{
                minWidth: '280px',
              }}
            >
              <PrimaryButton
                onClick={onClickToSearchPage}
              >
                検索
              </PrimaryButton>              
            </Box>
          </Grid>
        </Grid>
      </Box>
    </>
  )
}

export default SearchSection;