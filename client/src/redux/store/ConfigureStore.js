import { applyMiddleware, combineReducers, createStore, compose } from "redux";
import { persistStore, persistReducer } from "redux-persist";
import { encryptTransform } from "redux-persist-transform-encrypt";
import storage from "redux-persist/lib/storage";
import { thunk } from "redux-thunk";
import logger from "redux-logger";
import utilsReducer from "../reducers/utils.reducer";

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

const middleware =
  process.env.NODE_ENV === "development"
    ? applyMiddleware(thunk, logger)
    : applyMiddleware(thunk);

const encryptor = encryptTransform({
  secretKey: "dummy-secret",
  onError: function (error) {
    // Handle the error.
  },
});

const persistConfig = {
  key: "root",
  storage,
  transforms: [encryptor],
  blacklist: [],
};

const store = createStore(
  persistReducer(
    persistConfig,
    combineReducers({
      utils: utilsReducer,
    })
  ),
  composeEnhancers(middleware)
);

const persistor = persistStore(store);

export { store, persistor };
