import { VFC } from 'react';
import Select from 'react-select'

interface Tag {
  id: number;
  value: string;
  label: string;
}

type Props = {
  isMulti: boolean
  options: Tag[]
  onChange: (event:any) => void
  value?: Tag[]
}


export const TagSelction: VFC<Props> = (props) => {
  const {isMulti, onChange, options, value} = props

  if (isMulti) {
    return(
      <Select 
        placeholder="分野を選択してください"
        isMulti
        options={options}
        onChange={onChange}
        value={value}
      />
    )
  } else {
    return(
      <Select 
      placeholder="分野を選択してください"
      options={options}
      onChange={onChange}
      value={value}
    />
    )
  }
}

export default TagSelction; 