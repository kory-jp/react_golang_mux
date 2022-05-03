import React, { FC } from "react";

type Props = {
  text: string
}

export const TextFormat: FC<Props> = (props) => {
  const {text} = props
  const texts = text ? text.split(/(\n)/).map((item, index) => {
    return (
      <React.Fragment key={index}>
        {
          item.match(/\n/) ? <br/> : item
        }
      </React.Fragment>
    )
  }) : null
  return <>{texts}</>
}

export default TextFormat;