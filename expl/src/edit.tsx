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

export const CategoryEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.JeppCategoryInterface>({
            redirect: false,
            metaData: {
                fields: [
                    "name",
                    "id",
                    {
                        clues: [
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

    const [cluesCursors, setCluesCursors] = useState<Cursors>({});

    const { selectProps: cluesSelectProps } =
        RA.useSelect<Interfaces.JeppClueInterface>({
            resource: "Clue",
            optionLabel: "id",
            optionValue: "id",
            metaData: {
                cursors: cluesCursors,
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
                    <Action.CategoryShowAction recordItemIDs={id ? [id] : []} />
                </>
            )}
        >
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    label="Clues"
                    name={["clueIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...cluesSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Edit>
    );
};

export const ClueEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.JeppClueInterface>({
            redirect: false,
            metaData: {
                fields: [
                    "question",
                    "answer",
                    "categoryID",
                    "gameID",
                    "id",
                    {
                        category: ["id"],
                    },
                    {
                        game: ["id"],
                    },
                ],
            },
        });

    const [categoryCursors, setCategoryCursors] = useState<Cursors>({});

    const { selectProps: categorySelectProps } =
        RA.useSelect<Interfaces.JeppCategoryInterface>({
            resource: "Category",
            optionLabel: "name",
            optionValue: "id",
            metaData: {
                cursors: categoryCursors,
                fields: ["id", "name"],
            },
            onSearch: (value: string) => [
                {
                    field: "name",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [gameCursors, setGameCursors] = useState<Cursors>({});

    const { selectProps: gameSelectProps } =
        RA.useSelect<Interfaces.JeppGameInterface>({
            resource: "Game",
            optionLabel: "show",
            optionValue: "id",
            metaData: {
                cursors: gameCursors,
                fields: ["id", "show"],
            },
            onSearch: (value: string) => [
                {
                    field: "show",
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
                    <Action.ClueShowAction recordItemIDs={id ? [id] : []} />
                </>
            )}
        >
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    label="Category"
                    name="categoryID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...categorySelectProps} mode={undefined} />
                </Antd.Form.Item>
                <Antd.Form.Item
                    label="Game"
                    name="gameID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...gameSelectProps} mode={undefined} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Edit>
    );
};

export const GameEdit: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.JeppGameInterface>({
            redirect: false,
            metaData: {
                fields: [
                    "show",
                    "airDate",
                    "tapeDate",
                    "seasonID",
                    "id",
                    {
                        season: ["id"],
                    },
                    {
                        clues: [
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

    const [seasonCursors, setSeasonCursors] = useState<Cursors>({});

    const { selectProps: seasonSelectProps } =
        RA.useSelect<Interfaces.JeppSeasonInterface>({
            resource: "Season",
            optionLabel: "number",
            optionValue: "id",
            metaData: {
                cursors: seasonCursors,
                fields: ["id", "number"],
            },
            onSearch: (value: string) => [
                {
                    field: "number",
                    operator: "contains",
                    value,
                },
            ],
        });
    const [cluesCursors, setCluesCursors] = useState<Cursors>({});

    const { selectProps: cluesSelectProps } =
        RA.useSelect<Interfaces.JeppClueInterface>({
            resource: "Clue",
            optionLabel: "id",
            optionValue: "id",
            metaData: {
                cursors: cluesCursors,
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
                <Antd.Form.Item
                    label="Clues"
                    name={["clueIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...cluesSelectProps} mode={"multiple"} />
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
                    "startDate",
                    "endDate",
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
            optionLabel: "show",
            optionValue: "id",
            metaData: {
                cursors: gamesCursors,
                fields: ["id", "show"],
            },
            onSearch: (value: string) => [
                {
                    field: "show",
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
