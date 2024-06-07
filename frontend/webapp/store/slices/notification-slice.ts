import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Notification } from '@/types';

interface NotificationState {
  notifications: Notification[];
}

const initialState: NotificationState = {
  notifications: [],
};

export const notificationSlice = createSlice({
  name: 'notification',
  initialState,
  reducers: {
    addNotification: (
      state,
      action: PayloadAction<Omit<Notification, 'seen'>>
    ) => {
      state.notifications.push({
        ...action.payload,
        seen: false,
        isNew: true,
        time: new Date().toISOString(),
      });
    },
    markAsSeen: (state, action: PayloadAction<string>) => {
      const notification = state.notifications.find(
        (n) => n.id === action.payload
      );
      if (notification) {
        notification.seen = true;
      }
    },
    markAsOld: (state, action: PayloadAction<string>) => {
      const notification = state.notifications.find(
        (n) => n.id === action.payload
      );
      if (notification) {
        notification.isNew = false;
      }
    },
    removeNotification: (state, action: PayloadAction<string>) => {
      state.notifications = state.notifications.filter(
        (notification) => notification.id !== action.payload
      );
    },
  },
});

export const { addNotification, removeNotification, markAsSeen, markAsOld } =
  notificationSlice.actions;

export const notificationReducer = notificationSlice.reducer;
