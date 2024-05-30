'use client';
import React, { useEffect, useState } from 'react';
import theme from '@/styles/palette';
import { ACTION_ICONS } from '@/assets';
import { useActionState } from '@/hooks';
import { useSearchParams } from 'next/navigation';
import { ACTION, ACTIONS, ACTION_ITEM_DOCS_LINK } from '@/utils';
import {
  DeleteAction,
  DynamicActionForm,
  MultiCheckboxComponent,
} from '@/components';
import {
  KeyvalButton,
  KeyvalInput,
  KeyvalLink,
  KeyvalLoader,
  KeyvalSwitch,
  KeyvalText,
  KeyvalTextArea,
} from '@/design.system';
import {
  HeaderText,
  LoaderWrapper,
  DescriptionWrapper,
  CreateButtonWrapper,
  CreateActionWrapper,
  KeyvalInputWrapper,
  TextareaWrapper,
  SwitchWrapper,
  FormFieldsWrapper,
} from './styled';

const ACTION_ID = 'id';

export function EditActionContainer(): React.JSX.Element {
  const {
    actionState,
    onChangeActionState,
    upsertAction,
    buildActionData,
    onDeleteAction,
  } = useActionState();

  const {
    actionName,
    actionNote,
    actionData,
    selectedMonitors,
    disabled,
    type,
  } = actionState;

  const search = useSearchParams();

  useEffect(() => {
    const actionId = search.get(ACTION_ID);
    if (!actionId) return;
    buildActionData(actionId);
  }, [search]);

  if (!actionState || !type)
    return (
      <LoaderWrapper>
        <KeyvalLoader />
      </LoaderWrapper>
    );
  const ActionIcon = ACTION_ICONS[type];
  return (
    <CreateActionWrapper>
      <HeaderText>
        <ActionIcon style={{ width: 34, height: 34 }} />
        <KeyvalText size={18} weight={700}>
          {ACTIONS[type].TITLE}
        </KeyvalText>
        <SwitchWrapper disabled={disabled}>
          <KeyvalSwitch
            toggle={!disabled}
            handleToggleChange={() =>
              onChangeActionState('disabled', !disabled)
            }
            label={disabled ? ACTION.DISABLE : ACTION.APPLIED}
          />
        </SwitchWrapper>
      </HeaderText>
      <DescriptionWrapper>
        <KeyvalText size={14}>{ACTIONS[type].DESCRIPTION}</KeyvalText>
        <KeyvalLink
          value={ACTION.LINK_TO_DOCS}
          fontSize={14}
          onClick={() =>
            window.open(
              `${ACTION_ITEM_DOCS_LINK}/${type.toLowerCase()}`,
              '_blank'
            )
          }
        />
      </DescriptionWrapper>
      <FormFieldsWrapper disabled={disabled}>
        <MultiCheckboxComponent
          title={ACTIONS.MONITORS_TITLE}
          checkboxes={selectedMonitors}
          onSelectionChange={(newMonitors) =>
            onChangeActionState('selectedMonitors', newMonitors)
          }
        />
        <KeyvalInputWrapper>
          <KeyvalInput
            label={ACTIONS.ACTION_NAME}
            value={actionName}
            onChange={(name) => onChangeActionState('actionName', name)}
          />
        </KeyvalInputWrapper>
        <DynamicActionForm
          type={type}
          data={actionData}
          onChange={onChangeActionState}
        />
        <TextareaWrapper>
          <KeyvalTextArea
            label={ACTIONS.ACTION_NOTE}
            value={actionNote}
            placeholder={ACTIONS.NOTE_PLACEHOLDER}
            onChange={(e) => onChangeActionState('actionNote', e.target.value)}
          />
        </TextareaWrapper>
        <CreateButtonWrapper>
          <KeyvalButton onClick={upsertAction} disabled={!actionData}>
            <KeyvalText weight={600} color={theme.text.dark_button} size={14}>
              {ACTIONS.UPDATE_ACTION}
            </KeyvalText>
          </KeyvalButton>
        </CreateButtonWrapper>
        <DeleteAction onDelete={onDeleteAction} name={actionName} type={type} />
      </FormFieldsWrapper>
    </CreateActionWrapper>
  );
}
