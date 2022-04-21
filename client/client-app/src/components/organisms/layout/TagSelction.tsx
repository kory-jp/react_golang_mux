import { VFC } from 'react';
import Select from 'react-select'

interface Tag {
  id: number;
  value: string;
  label: string;
}

type Props = {
  options: Tag[]
  onChange: (event:any) => void
  value?: Tag[]
}


export const TagSelction: VFC<Props> = (props) => {
  const { onChange, options, value} = props

  return(
    <Select 
      placeholder="分野を選択してください"
      isMulti
      name="tag"
      options={options}
      onChange={onChange}
      value={value}
    />
  )
}

export default TagSelction; 