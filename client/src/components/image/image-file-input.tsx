import type { ChangeEvent, FC } from "react";

type Properties = {
  setSelectedFile: (file: File) => void;
};

const ImageFileInput: FC<Properties> = ({ setSelectedFile }) => {
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.item(0);
    if (file) {
      setSelectedFile(file);
    }
  };

  return (
    <div>
      <label htmlFor="image">Choose an image:</label>
      <input
        type="file"
        id="file"
        name="file"
        accept="image/png"
        onChange={(event) => handleChange(event)}
      />
    </div>
  );
};

export default ImageFileInput;
