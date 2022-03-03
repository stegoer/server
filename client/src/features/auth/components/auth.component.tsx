import { isFormType } from "@features/auth/auth.types";
import { DEFAULT_FORM_TYPE } from "@features/auth/components/auth-form/auth-form.constants";
import AuthForm from "@features/auth/components/auth-form/auth.form";
import AuthFormSkeleton from "@features/auth/components/auth-form/skeleton/auth-form.skeleton";

import { Title } from "@mantine/core";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

import type { FormType } from "@features/auth/auth.types";

const AuthComponent = (): JSX.Element => {
  const router = useRouter();
  const { form: formQuery } = router.query;
  const [formType, setFormType] = useState<FormType>(DEFAULT_FORM_TYPE);

  // once router is ready set formType if formQuery contains a valid value
  useEffect(() => {
    if (
      router.isReady &&
      typeof formQuery === `string` &&
      isFormType(formQuery)
    ) {
      setFormType(formQuery);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [router.isReady]);

  // whenever `formType` is changed update query params accordingly
  useEffect(() => {
    if (router.isReady) {
      void router.push(
        {
          pathname: router.pathname,
          query: { ...router.query, form: formType },
        },
        `${router.pathname}?form=${formType}`,
        { shallow: true },
      );
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [formType]);

  return router.isReady ? (
    <AuthForm
      formType={formType}
      toggleFormType={() =>
        setFormType((previous) =>
          previous === `login` ? `register` : `login`,
        )
      }
    />
  ) : (
    <>
      <Title>Account</Title>
      <AuthFormSkeleton />
    </>
  );
};

export default AuthComponent;
