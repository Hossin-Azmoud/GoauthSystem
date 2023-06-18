import { useEffect, useState } from 'react';
import { generateID } from '../core/core.jsx';

const Input = ({
  Label,
  setter,
  validate=()=> [],
  ...Rest
}) => {
  const Id = generateID();
  const [Err, setErr] = useState("");

  useEffect(() => {

  }, [val])

  return (
    <label htmlFor={Id} className="flex flex-col my-2 w-[90%] sm:w-[400px]">
      { Label }
      <input value={ val } onChange={ setter } id={ Id } className="my-2 p-2 border outline-none rounded" {...Rest}/> 
      {
        (val) ? (
          <span className="text-red"> { validate(val) } </span>
        ) : ()
      }
    </label>
  )
}

export { Input };


