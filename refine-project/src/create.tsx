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
export const CategoryCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.JeppCategoryInterface>();

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

    return (
        <RA.Create
            saveButtonProps={saveButtonProps}
            headerButtons={() => <></>}
        >
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="name"
                    label="Name"
                    rules={[{ required: true }]}
                >
                    <View.JeppStringViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Clues"
                    name={["clueIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...cluesSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const ClueCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.JeppClueInterface>();

    const [categoryCursors, setCategoryCursors] = useState<Cursors>({});

    const { selectProps: categorySelectProps } =
        RA.useSelect<Interfaces.JeppCategoryInterface>({
            resource: "Category",
            optionLabel: "id",
            optionValue: "id",
            metaData: {
                cursors: categoryCursors,
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

    return (
        <RA.Create
            saveButtonProps={saveButtonProps}
            headerButtons={() => <></>}
        >
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="question"
                    label="Question"
                    rules={[{ required: true }]}
                >
                    <View.JeppStringViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    name="answer"
                    label="Answer"
                    rules={[{ required: true }]}
                >
                    <View.JeppStringViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    name="categoryID"
                    label="Category Id"
                    rules={[{ required: false }]}
                >
                    <View.JeppNumberViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    name="gameID"
                    label="Game Id"
                    rules={[{ required: false }]}
                >
                    <View.JeppNumberViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Category"
                    name="categoryID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...categorySelectProps} mode={undefined} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const GameCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.JeppGameInterface>();

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

    return (
        <RA.Create
            saveButtonProps={saveButtonProps}
            headerButtons={() => <></>}
        >
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="show"
                    label="Show"
                    rules={[{ required: true }]}
                >
                    <View.JeppNumberViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    name="airdate"
                    label="AirDate"
                    rules={[{ required: true }]}
                    getValueProps={(value) => ({
                        value: value ? dayjs(value) : "",
                    })}
                >
                    <View.JeppDateViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    name="tapedate"
                    label="TapeDate"
                    rules={[{ required: true }]}
                    getValueProps={(value) => ({
                        value: value ? dayjs(value) : "",
                    })}
                >
                    <View.JeppDateViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    name="seasonID"
                    label="Season Id"
                    rules={[{ required: false }]}
                >
                    <View.JeppNumberViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Season"
                    name="seasonID"
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...seasonSelectProps} mode={undefined} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};
export const SeasonCreate: React.FC = () => {
    const { formProps, saveButtonProps, queryResult } =
        RA.useForm<Interfaces.JeppSeasonInterface>();

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

    return (
        <RA.Create
            saveButtonProps={saveButtonProps}
            headerButtons={() => <></>}
        >
            <Antd.Form {...formProps} layout="vertical">
                <Antd.Form.Item
                    name="number"
                    label="Number"
                    rules={[{ required: true }]}
                >
                    <View.JeppNumberViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    name="startdate"
                    label="StartDate"
                    rules={[{ required: true }]}
                    getValueProps={(value) => ({
                        value: value ? dayjs(value) : "",
                    })}
                >
                    <View.JeppDateViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    name="enddate"
                    label="EndDate"
                    rules={[{ required: true }]}
                    getValueProps={(value) => ({
                        value: value ? dayjs(value) : "",
                    })}
                >
                    <View.JeppDateViewOnForm />
                </Antd.Form.Item>

                <Antd.Form.Item
                    label="Games"
                    name={["gameIDs"]}
                    rules={[{ required: false }]}
                >
                    <Antd.Select {...gamesSelectProps} mode={"multiple"} />
                </Antd.Form.Item>
            </Antd.Form>
        </RA.Create>
    );
};