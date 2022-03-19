import { CHANNEL_SWITCH_STYLES } from "@features/images/images.constants";
import { Channel } from "@graphql/generated/codegen.generated";

import { Switch } from "@mantine/core";
import { useEffect, useState } from "react";

import type {
  ChannelSwitchStateType,
  ChannelSwitchType,
} from "@features/images/images.types";
import type { UseForm } from "@mantine/hooks/lib/use-form/use-form";

export type ChannelSwitchesProps<T extends { channel?: Channel }> = {
  form: UseForm<T>;
  disabled: boolean;
};

const ChannelSwitches = <T extends { channel?: Channel }>({
  form,
  disabled,
}: ChannelSwitchesProps<T>): JSX.Element => {
  const [redChecked, setRedChecked] = useState(true);
  const [greenChecked, setGreenChecked] = useState(true);
  const [blueChecked, setBlueChecked] = useState(true);

  const switchStates: ChannelSwitchStateType[] = [
    { checked: redChecked, setChecked: setRedChecked },
    { checked: greenChecked, setChecked: setGreenChecked },
    { checked: blueChecked, setChecked: setBlueChecked },
  ];
  const switches: ChannelSwitchType[] = CHANNEL_SWITCH_STYLES.map(
    (style, index) => {
      return { style, state: switchStates[index] };
    },
  );

  useEffect(() => {
    let updatedChannel: Channel | undefined;
    if (redChecked && greenChecked && blueChecked) {
      updatedChannel = Channel.RedGreenBlue;
    } else if (redChecked && greenChecked && !blueChecked) {
      updatedChannel = Channel.RedGreen;
    } else if (redChecked && !greenChecked && blueChecked) {
      updatedChannel = Channel.RedBlue;
    } else if (redChecked && !greenChecked && !blueChecked) {
      updatedChannel = Channel.Red;
    } else if (!redChecked && greenChecked && blueChecked) {
      updatedChannel = Channel.GreenBlue;
    } else if (!redChecked && greenChecked && !blueChecked) {
      updatedChannel = Channel.Green;
    } else if (!redChecked && !greenChecked && blueChecked) {
      updatedChannel = Channel.Blue;
    }
    form.setFieldValue(`channel`, updatedChannel);
    form.validateField(`channel`);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [blueChecked, greenChecked, redChecked]);

  return (
    <>
      {switches.map((channelSwitch, index) => (
        <Switch
          key={index}
          label={channelSwitch.style.label}
          color={channelSwitch.style.color}
          checked={channelSwitch.state.checked}
          onChange={(event) =>
            channelSwitch.state.setChecked(event.currentTarget.checked)
          }
          onBlur={() => form.validateField(`channel`)}
          disabled={disabled}
        />
      ))}
    </>
  );
};

export default ChannelSwitches;
