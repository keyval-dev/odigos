import { ROUTES } from '@/utils';
import { useState } from 'react';
import { useMutation } from 'react-query';
import { useActions } from './useActions';
import { useRouter } from 'next/navigation';
import { ActionData, ActionItem, ActionsType } from '@/types';
import { putAction, setAction, deleteAction } from '@/services';

interface Monitor {
  id: string;
  label: string;
  checked: boolean;
}

interface ActionState {
  id?: string;
  actionName: string;
  actionNote: string;
  actionData: any;
  selectedMonitors: Monitor[];
  disabled: boolean;
  type: string;
}

const DEFAULT_MONITORS: Monitor[] = [
  { id: '1', label: 'Logs', checked: true },
  { id: '2', label: 'Metrics', checked: true },
  { id: '3', label: 'Traces', checked: true },
];

export function useActionState() {
  const [actionState, setActionState] = useState<ActionState>({
    actionName: '',
    actionNote: '',
    actionData: null,
    selectedMonitors: DEFAULT_MONITORS,
    disabled: false,
    type: '',
  });

  const router = useRouter();
  const { getActionById } = useActions();

  const { mutateAsync: createAction } = useMutation((body: ActionItem) =>
    setAction(body, actionState.type)
  );
  const { mutateAsync: updateAction } = useMutation((body: ActionItem) =>
    putAction(actionState?.id, body, actionState.type)
  );

  const { mutateAsync: deleteActionMutation } = useMutation((id: string) =>
    deleteAction(id, actionState.type)
  );

  async function onSuccess() {
    router.push(ROUTES.ACTIONS);
  }

  function onChangeActionState(key: string, value: any) {
    setActionState((prevState) => ({
      ...prevState,
      [key]: value,
    }));
    if (key === 'disabled') upsertAction(false);
  }

  async function buildActionData(actionId: string) {
    const action = await getActionById(actionId);

    const actionState = {
      id: action?.id,
      actionName: action?.spec?.actionName || '',
      actionNote: action?.spec?.notes || '',
      type: action?.type || '',
      actionData: getActionDataByType(action),
      selectedMonitors: DEFAULT_MONITORS.map((monitor) => ({
        ...monitor,

        checked: !!action?.spec?.signals.includes(monitor.label.toUpperCase()),
      })),
      disabled: action?.spec?.disabled || false,
    };

    setActionState(actionState);
  }

  async function upsertAction(callback: boolean = true) {
    const {
      actionName,
      actionNote,
      actionData,
      selectedMonitors,
      disabled,
      type,
    } = actionState;

    const signals = selectedMonitors
      .filter((monitor) => monitor.checked)
      .map((monitor) => monitor.label.toUpperCase());

    const filteredActionData = filterEmptyActionDataFieldsByType(
      type,
      actionData
    );

    const action = {
      actionName,
      notes: actionNote,
      signals,
      ...filteredActionData,
      disabled: callback ? disabled : !disabled,
    };

    try {
      if (actionState?.id) {
        await updateAction(action);
      } else {
        delete action.disabled;
        await createAction(action);
      }
      callback && onSuccess();
    } catch (error) {
      console.error({ error });
    }
  }

  function onDeleteAction() {
    try {
      if (actionState?.id) {
        deleteActionMutation(actionState.id);
        onSuccess();
      }
    } catch (error) {}
  }

  return {
    actionState,
    onChangeActionState,
    upsertAction,
    buildActionData,
    onDeleteAction,
  };
}

function filterEmptyActionDataFieldsByType(type: string, data: any) {
  switch (type) {
    case ActionsType.ADD_CLUSTER_INFO:
      return {
        clusterAttributes: data.clusterAttributes.filter(
          (attr: any) =>
            attr.attributeStringValue !== '' && attr.attributeName !== ''
        ),
      };
    case ActionsType.DELETE_ATTRIBUTES:
      return {
        attributeNamesToDelete: data.attributeNamesToDelete.filter(
          (attr: string) => attr !== ''
        ),
      };

    default:
      return data;
  }
}

function getActionDataByType(action: ActionData | undefined) {
  if (!action) return {};
  switch (action.type) {
    case ActionsType.ADD_CLUSTER_INFO:
      return {
        clusterAttributes: action.spec.clusterAttributes.map((attr, index) => ({
          attributeName: attr.attributeName,
          attributeStringValue: attr.attributeStringValue,
          id: index,
        })),
      };
    case ActionsType.DELETE_ATTRIBUTES:
      return {
        attributeNamesToDelete: action.spec.attributeNamesToDelete,
      };
    default:
      return {};
  }
}
