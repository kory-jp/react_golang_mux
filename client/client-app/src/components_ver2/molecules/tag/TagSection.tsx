import { Grid } from "@mui/material";
import { VFC } from "react";
import { Tags } from "../../../reducks/tags/types";
import { PrimaryChip } from "../../atoms/chip/PrimaryChip";

type Props = {
  tags: Tags | null,
  onClickToSearchTagTodo: (id: number) => void,
}

export const TagSection: VFC<Props> = (props) => {
  const {tags, onClickToSearchTagTodo} = props
  return(
    <>
      <Grid
        container
        sx={{
          minHeight: '40px',
        }}
      >
        {
          tags != null && (
            <>
              {
                tags.map(tag => (
                  <Grid
                    key={tag.id}
                  >
                    <PrimaryChip 
                      label={tag.label}
                      onClick={() => onClickToSearchTagTodo(tag.id)}
                    />                      
                  </Grid>
                ))
              }                
            </>
          )
        }
      </Grid>      
    </>
  )
}

export default TagSection;