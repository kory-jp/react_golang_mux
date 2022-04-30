import { FC } from "react";
import IndexSection from "../../organisms/toods/IndexSection";
import DefaultTemplate from "./DefaultTemplate";

export const IndexTodosTemplate: FC = () => {
  return(
    <>
      <DefaultTemplate>
        <IndexSection />
      </DefaultTemplate>
    </>
  )
}

export default IndexTodosTemplate