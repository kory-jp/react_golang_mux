import { FC } from "react";
import IndexSection from "../../organisms/toods/IndexSection";
import DefaultIndexTemplate from "./DefaultIndexTemplate";

export const IndexTodosTemplate: FC = () => {
  return(
    <>
      <DefaultIndexTemplate>
        <IndexSection />
      </DefaultIndexTemplate>
    </>
  )
}

export default IndexTodosTemplate