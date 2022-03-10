import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";
import type { ChangeEvent } from "react";

type Props<T extends { file?: File }> = {
  form: UseForm<T>;
  disabled: boolean;
};

const ImageFileInput = <T extends { file?: File }>({
  form,
  disabled,
}: Props<T>): JSX.Element => {
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.item(0) ?? undefined;
    form.setFieldValue(`file`, file);
  };

  return (
    <>
      <label htmlFor="image">Choose an image:</label>
      <input
        type="file"
        id="file"
        name="file"
        accept="image/png"
        disabled={disabled}
        onChange={(event) => handleChange(event)}
        onBlur={() => form.validateField(`file`)}
      />
    </>
  );
};

export default ImageFileInput;
