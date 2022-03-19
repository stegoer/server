import { LSB_USED_MARK, LSB_USED_MAX } from "@features/images/images.constants";

import { Slider } from "@mantine/core";

import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

export type LSBUsedSliderProps<T extends { lsbUsed: number }> = {
  form: UseForm<T>;
};

const MARKS = [
  { value: LSB_USED_MARK, label: `1` },
  { value: LSB_USED_MARK * 2, label: `2` },
  { value: LSB_USED_MARK * 3, label: `3` },
  { value: LSB_USED_MARK * 4, label: `4` },
  { value: LSB_USED_MARK * 5, label: `5` },
  { value: LSB_USED_MARK * 6, label: `6` },
  { value: LSB_USED_MARK * 7, label: `7` },
  { value: LSB_USED_MARK * 8, label: `8` },
];

const LSBUsedSlider = <T extends { lsbUsed: number }>({
  form,
}: LSBUsedSliderProps<T>): JSX.Element => {
  return (
    <Slider
      style={{ marginTop: `20px`, marginBottom: `30px` }}
      label={(value) => MARKS.find((mark) => mark.value === value)?.label}
      defaultValue={LSB_USED_MARK}
      min={LSB_USED_MARK}
      max={LSB_USED_MARK * LSB_USED_MAX}
      step={LSB_USED_MARK}
      marks={MARKS}
      labelAlwaysOn
      {...form.getInputProps(`lsbUsed`)}
    />
  );
};

export default LSBUsedSlider;
