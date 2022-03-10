import SubmitButton from "@components/buttons/submit.button";
import ErrorText from "@components/errors/error.text";
import ImagesFormInput from "@features/images/components/images-form/images-form.input";
import useImagesForm from "@hooks/images-form.hook";
import { capitalize } from "@utils/format.utils";

import { LoadingOverlay } from "@mantine/core";
import { useEffect } from "react";

import type { FormType, UseFormType } from "@features/images/images.types";
import type { ReactNode, PropsWithChildren } from "react";

type Props = PropsWithChildren<{
  formType: FormType;
  loading: boolean;
  onSubmit(values: UseFormType[`values`]): void;
  error?: ReactNode;
  setError(error?: ReactNode): void;
}>;

const ImagesFormComponent = ({
  formType,
  loading,
  onSubmit,
  error,
  setError,
  children,
}: Props): JSX.Element => {
  const form = useImagesForm(formType);

  useEffect(() => {
    const updatedError =
      (form.errors.message ||
        form.errors.lsbUsed ||
        form.errors.channel ||
        form.errors.file) ??
      undefined;
    if (updatedError) {
      setError(updatedError);
    }
  }, [
    form.errors.channel,
    form.errors.file,
    form.errors.lsbUsed,
    form.errors.message,
    setError,
  ]);

  return (
    <>
      <form onSubmit={form.onSubmit(onSubmit)}>
        <LoadingOverlay visible={loading} />

        <ImagesFormInput form={form} formType={formType} disabled={loading} />

        {error && <ErrorText error={error} />}

        <SubmitButton disabled={loading}>{capitalize(formType)}</SubmitButton>
      </form>
      {children}
    </>
  );
};

export default ImagesFormComponent;
