import { FC } from "react";
import IndexTCSection from "../../organisms/taskCards/IndexTCSection";
import ShowSection from "../../organisms/toods/ShowSection";
import DefaultTemplate from "./DefaultTemplate";

export const  ShowTodoTemplate: FC = () => {
  return (
    <>
      <DefaultTemplate>
        <ShowSection />
        <IndexTCSection />
      </DefaultTemplate>
    </>
  )
}

export default  ShowTodoTemplate;