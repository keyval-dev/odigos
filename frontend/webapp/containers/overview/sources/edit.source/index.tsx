'use client';
import { ManageSourceHeader } from '@/components/overview/sources/manage.source.header/manage.source.header';
import {
  ACTION,
  LANGUAGES_LOGOS,
  NOTIFICATION,
  OVERVIEW,
  ROUTES,
  SETUP,
  getMainContainerLanguage,
} from '@/utils';
import { useRouter, useSearchParams } from 'next/navigation';
import React, { useEffect, useState } from 'react';
import { useMutation } from 'react-query';
import {
  ManageSourcePageContainer,
  BackButtonWrapper,
  FieldWrapper,
  SaveSourceButtonWrapper,
} from './styled';
import { Back } from '@/assets/icons/overview';
import {
  KeyvalButton,
  KeyvalInput,
  KeyvalLoader,
  KeyvalText,
} from '@/design.system';
import { DeleteSource } from '@/components/overview';
import { deleteSource, getSource, patchSources } from '@/services/sources';
import { useKeyDown, useNotification } from '@/hooks';
import theme from '@/styles/palette';
import { ManagedSource } from '@/types/sources';
import ConditionAlert from '@/design.system/condition.alert';

const NAME = 'name';
const KIND = 'kind';
const NAMESPACE = 'namespace';

const exampleConditions = [
  {
    type: 'AppliedInstrumentationDevice',
    status: 'True',
    message: 'Successfully applied instrumentation device to pod template',
    last_transition_time: '2024-05-21 11:44:53 +0300 IDT',
  },
  {
    type: 'FailedToLoad',
    status: 'False',
    message: 'offset not found: net/http.response:status (1.22.3)',
    last_transition_time: '2024-05-21 11:45:00 +0300 IDT',
  },
];

export function EditSourceForm() {
  const [inputValue, setInputValue] = useState('');
  const [currentSource, setCurrentSource] = useState<ManagedSource>();

  const searchParams = useSearchParams();
  const router = useRouter();
  const { show, Notification } = useNotification();

  const { mutate: handleDeleteSource } = useMutation(() =>
    deleteSource(
      currentSource?.namespace || '',
      currentSource?.kind || '',
      currentSource?.name || ''
    )
  );

  const { mutate: editSource } = useMutation(() =>
    patchSources(
      currentSource?.namespace || '',
      currentSource?.kind || '',
      currentSource?.name || '',
      { reported_name: inputValue }
    )
  );
  useEffect(() => {
    onPageLoad();
  }, [searchParams]);

  useEffect(() => {
    setInputValue(currentSource?.reported_name || '');
  }, [currentSource]);

  useKeyDown('Enter', handleKeyPress);

  function handleKeyPress(e: any) {
    onSaveClick();
  }

  async function onPageLoad() {
    const name = searchParams.get(NAME) || '';
    const kind = searchParams.get(KIND) || '';
    const namespace = searchParams.get(NAMESPACE) || '';

    const currentSource = await getSource(namespace, kind, name);
    setCurrentSource(currentSource);
  }
  function onError({ response }) {
    const message = response?.data?.message;
    show({
      type: NOTIFICATION.ERROR,
      message,
    });
  }

  function onSaveClick() {
    editSource(undefined, {
      onError,
      onSuccess: () => router.push(`${ROUTES.SOURCES}?status=updated`),
    });
  }

  function onSourceDelete() {
    handleDeleteSource(undefined, {
      onSuccess: () => router.push(`${ROUTES.SOURCES}?status=deleted`),
      onError,
    });
  }

  if (!currentSource) {
    return <KeyvalLoader />;
  }

  return (
    <ManageSourcePageContainer>
      <BackButtonWrapper onClick={() => router.back()}>
        <Back width={14} />
        <KeyvalText size={14}>{SETUP.BACK}</KeyvalText>
      </BackButtonWrapper>
      {currentSource && <ManageSourceHeader source={currentSource} />}
      <div style={{ display: 'flex', gap: 60 }}>
        <div>
          <FieldWrapper>
            <KeyvalInput
              label={OVERVIEW.REPORTED_NAME}
              value={inputValue}
              onChange={(e) => setInputValue(e)}
            />
          </FieldWrapper>
          <SaveSourceButtonWrapper>
            <KeyvalButton disabled={!inputValue} onClick={onSaveClick}>
              <KeyvalText color={theme.colors.dark_blue} size={14} weight={600}>
                {ACTION.SAVE}
              </KeyvalText>
            </KeyvalButton>
          </SaveSourceButtonWrapper>
          <DeleteSource
            onDelete={onSourceDelete}
            name={currentSource?.name}
            image_url={
              LANGUAGES_LOGOS[
                getMainContainerLanguage(
                  currentSource?.instrumented_application_details.languages ||
                    undefined
                )
              ]
            }
          />
        </div>
        <ConditionAlert conditions={currentSource.conditions} />
      </div>

      <Notification />
    </ManageSourcePageContainer>
  );
}
