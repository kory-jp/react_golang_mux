import { FC } from "react";
import SearchIndexSection from "../../organisms/toods/SearchIndexSection";
import DefaultIndexTemplate from "./DefaultIndexTemplate";

export const SearchTodosTemplate: FC = () => {
  return(
    <>
      <DefaultIndexTemplate>
        <SearchIndexSection />
      </DefaultIndexTemplate>
    </>
  )
}

export default SearchTodosTemplate