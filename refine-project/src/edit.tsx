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

import { useState } from "react";
import * as RA from "@refinedev/antd";
import * as Antd from "antd";
import * as Interfaces from "./typedefs";
import { Cursors } from "./data-provider";
import dayjs from "dayjs";
import CodeEditor from "@uiw/react-textarea-code-editor";
import * as View from "./view";
import * as Action from "./action";
import * as Custom from "./custom";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";

export const GameEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.JeppGameInterface>({
            redirect: false,
            metaData: {
                fields: [
                    "show",
                    "airdate",
                    "tapedate",
                    "id",
                    {
                        season: ["id"],
                    },
                ],
            },
        });

    const [seasonCursors, setSeasonCursors] = useState<Cursors>({});

    const { selectProps: seasonSelectProps } =
        RA.useSelect<Interfaces.JeppSeasonInterface>({
            resource: "Season",
            optionLabel: "id",
            optionValue: "id",
            metaData: {
                cursors: seasonCursors,
                fields: ["id", "id"],
            },
            onSearch: (value: string) => [
                {
                    field: "id",
                    operator: "contains",
                    value,
                },
            ],
        });

    const id = queryResult?.data?.data.id;
    return (
        <RA.Edit
            saveButtonProps={saveButtonProps}
            headerButtons={() => (
                <>
                    <Action.GameShowAction recordItemIDs={id ? [id] : []} />
                </>
            )}
        >
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    label="Season"
                    name="seasonID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...seasonSelectProps} mode={undefined} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Edit>
    );
};

export const SeasonEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.JeppSeasonInterface>({
            redirect: false,
            metaData: {
                fields: [
                    "number",
                    "startdate",
                    "enddate",
                    "id",
                    {
                        games: [
                            {
                                edges: [
                                    {
                                        node: ["id"],
                                    },
                                ],
                            },
                        ],
                    },
                ],
            },
        });

    const [gamesCursors, setGamesCursors] = useState<Cursors>({});

    const { selectProps: gamesSelectProps } =
        RA.useSelect<Interfaces.JeppGameInterface>({
            resource: "Game",
            optionLabel: "id",
            optionValue: "id",
            metaData: {
                cursors: gamesCursors,
                fields: ["id", "id"],
            },
            onSearch: (value: string) => [
                {
                    field: "id",
                    operator: "contains",
                    value,
                },
            ],
        });

    const id = queryResult?.data?.data.id;
    return (
        <RA.Edit
            saveButtonProps={saveButtonProps}
            headerButtons={() => (
                <>
                    <Action.SeasonShowAction recordItemIDs={id ? [id] : []} />
                </>
            )}
        >
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    label="Games"
                    name={["gameIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...gamesSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Edit>
    );
};
