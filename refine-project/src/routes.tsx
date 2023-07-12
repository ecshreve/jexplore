// Code generated by EntKit. DO NOT EDIT.
// ---------------------------------------------------------
//
// Copyright (C) 2023 EntKit. All Rights Reserved.
//
// This code is part of the EntKit library and is generated
// automatically to ensure optimal functionality and maintainability.
// Any changes made directly to this file may be overwritten
// by future code generation, leading to unexpected behavior.
//
// Please refer to the EntKit documentation for instructions on
// how to modify the library, extend its functionality or contribute
// to the project: https://entkit.com
// ---------------------------------------------------------

import React from "react";

// Components imports
import * as List from "./list";
import * as Show from "./show";
import * as Create from "./create";
import * as Edit from "./edit";
import * as EdgesDiagram from "./edges-diagram";

import { Login } from "./login";
import { Header } from "./header";
import { Authenticated } from "@refinedev/core";
import { ErrorComponent, ThemedLayoutV2, ThemedTitleV2 } from "@refinedev/antd";
import { Route, Routes, Outlet } from "react-router-dom";
import {
    NavigateToResource,
    CatchAllNavigate,
} from "@refinedev/react-router-v6";

export const RoutesBundle: React.FC = () => {
    const loginUrl = window.environment.appPath + "login";
    return (
        <Routes>
            <Route
                path={window.environment.appPath}
                element={
                    <Authenticated
                        fallback={<CatchAllNavigate to={loginUrl} />}
                    >
                        <ThemedLayoutV2
                            Header={Header}
                            Title={({ collapsed }) => (
                                <ThemedTitleV2
                                    // collapsed is a boolean value that indicates whether the <Sidebar> is collapsed or not
                                    collapsed={collapsed}
                                    // icon={collapsed ? <MySmallIcon /> : <MyLargeIcon />}
                                    text="EntKit"
                                />
                            )}
                        >
                            <Outlet />
                        </ThemedLayoutV2>
                    </Authenticated>
                }
            >
                <Route index element={<NavigateToResource resource="game" />} />

                <Route path="game">
                    <Route path="show/:id" element={<Show.GameMainShow />} />
                    <Route index path="" element={<List.GameList />} />
                    <Route path="create" element={<Create.GameCreate />} />
                    <Route path="edit/:id" element={<Edit.GameEdit />} />
                </Route>

                <Route path="season">
                    <Route path="show/:id" element={<Show.SeasonMainShow />} />
                    <Route index path="" element={<List.SeasonList />} />
                    <Route path="create" element={<Create.SeasonCreate />} />
                    <Route path="edit/:id" element={<Edit.SeasonEdit />} />
                </Route>
            </Route>

            <Route
                path={window.environment.appPath}
                element={
                    <Authenticated fallback={<Outlet />}>
                        <NavigateToResource resource="game" />
                    </Authenticated>
                }
            >
                <Route path="login" element={<Login />} />
            </Route>

            <Route
                path={window.environment.appPath}
                element={
                    <Authenticated>
                        <ThemedLayoutV2>
                            <Outlet />
                        </ThemedLayoutV2>
                    </Authenticated>
                }
            >
                <Route path="*" element={<ErrorComponent />} />
            </Route>
        </Routes>
    );
};
