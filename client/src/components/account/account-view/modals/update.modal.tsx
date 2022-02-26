import UpdateForm from "@components/forms/update-form/update.form";

import { Modal } from "@mantine/core";

import type { User } from "@graphql/generated/codegen.generated";
import type { Dispatch, FC, SetStateAction } from "react";

type Props = {
  user: User;
  opened: boolean;
  setOpened: Dispatch<SetStateAction<boolean>>;
};

const UpdateModal: FC<Props> = ({ user, opened, setOpened }) => {
  return (
    <Modal
      opened={opened}
      onClose={() => setOpened(false)}
      title={`Update ${user.username} account`}
    >
      <UpdateForm user={user} />
    </Modal>
  );
};

export default UpdateModal;
