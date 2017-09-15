module BootstrapLayout exposing (main)

import Html exposing (..)
import Html.Attributes exposing (href, class)
import Html.Events exposing (onClick)
import Navigation exposing (Location)
import UrlParser exposing ((</>))
import Json.Decode exposing (int, string, float, list, Decoder)
import Json.Decode.Pipeline exposing (decode, required, optional, hardcoded)
import Bootstrap.Navbar as Navbar
import Bootstrap.Grid as Grid
import Bootstrap.Grid.Col as Col
import Bootstrap.Card as Card
import Bootstrap.Button as Button
import Bootstrap.ListGroup as ListGroup
import Bootstrap.Modal as Modal
import Http
import Debug
import Time exposing (..)


main : Program Never Model Msg
main =
    -- Use updateDebug to view model changes and messages during development, and update
    -- for production
    Navigation.program UrlChange
        { view = view
        , update = update
        , subscriptions = subscriptions
        , init = init
        }


type alias Model =
    { page : Page
    , navState : Navbar.State
    , modalState : Modal.State
    , samples : List Sample
    }


type alias Sample =
    { serialNumber : String
    , airRegulatorValueLoc : Float
    , cycleCountValue : Int
    , mode : Int
    , postSprayValue : Float
    , preSprayValue : Float
    , programValue : Int
    , sprayValue : Float
    }


type Page
    = Home
    | Samples
    | GettingStarted
    | Modules
    | NotFound


init : Location -> ( Model, Cmd Msg )
init location =
    let
        ( navState, navCmd ) =
            Navbar.initialState NavMsg

        ( model, urlCmd ) =
            urlUpdate location { navState = navState, page = Home, modalState = Modal.hiddenState, samples = [] }
    in
        ( model, Cmd.batch [ urlCmd, navCmd, getSamples ] )


type Msg
    = UrlChange Location
    | NavMsg Navbar.State
    | ModalMsg Modal.State
    | UpdateSamples (Result Http.Error (List Sample))
    | RefreshSamples
    | Tick Time


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.batch
        [ Navbar.subscriptions model.navState NavMsg
        , every (10 * second) Tick
        ]


updateDebug : Msg -> Model -> ( Model, Cmd Msg )
updateDebug msg model =
    let
        _ =
            Debug.log "IN model" model

        _ =
            Debug.log "IN msg" msg

        ( modelUp, msgUp ) =
            update msg model

        _ =
            Debug.log "RET model" modelUp

        _ =
            Debug.log "RET msg" msgUp
    in
        ( modelUp, msgUp )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        UrlChange location ->
            urlUpdate location model

        NavMsg state ->
            ( { model | navState = state }
            , Cmd.none
            )

        ModalMsg state ->
            ( { model | modalState = state }
            , Cmd.none
            )

        UpdateSamples (Ok samples) ->
            ( { model | samples = samples }
            , Cmd.none
            )

        UpdateSamples (Err _) ->
            ( model, Cmd.none )

        RefreshSamples ->
            ( model, getSamples )

        Tick time ->
            ( model, getSamples )


getSamples : Cmd Msg
getSamples =
    let
        url =
            "/sample"

        request =
            Http.get url sampleListDecoder
    in
        Http.send UpdateSamples request


sampleListDecoder : Decoder (List Sample)
sampleListDecoder =
    list sampleDecoder


sampleDecoder : Decoder Sample
sampleDecoder =
    decode Sample
        |> required "serialNumber" string
        |> required "airRegulatorValueLoc" float
        |> required "cycleCountValue" int
        |> required "mode" int
        |> required "postSprayValue" float
        |> required "preSprayValue" float
        |> required "programValue" int
        |> required "sprayValue" float


urlUpdate : Navigation.Location -> Model -> ( Model, Cmd Msg )
urlUpdate location model =
    case decodeLoc location of
        Nothing ->
            ( { model | page = NotFound }, Cmd.none )

        Just route ->
            ( { model | page = route }, Cmd.none )


decodeLoc : Location -> Maybe Page
decodeLoc location =
    UrlParser.parseHash routeParser location


routeParser : UrlParser.Parser (Page -> a) a
routeParser =
    UrlParser.oneOf
        [ UrlParser.map Home UrlParser.top
        , UrlParser.map Samples (UrlParser.s "samples")
        , UrlParser.map GettingStarted (UrlParser.s "getting-started")
        , UrlParser.map Modules (UrlParser.s "modules")
        ]


view : Model -> Html Msg
view model =
    div []
        [ menu model
        , mainContent model
        , modal model
        ]


menu : Model -> Html Msg
menu model =
    Navbar.config NavMsg
        |> Navbar.withAnimation
        |> Navbar.container
        |> Navbar.brand [ href "#" ] [ text "Go/Elm Demo" ]
        |> Navbar.items
            [ Navbar.itemLink [ href "#samples" ] [ text "Samples" ]
            , Navbar.itemLink [ href "#getting-started" ] [ text "Getting started" ]
            , Navbar.itemLink [ href "#modules" ] [ text "Modules" ]
            ]
        |> Navbar.view model.navState


mainContent : Model -> Html Msg
mainContent model =
    Grid.container [] <|
        case model.page of
            Home ->
                pageHome model

            Samples ->
                pageSamples model

            GettingStarted ->
                pageGettingStarted model

            Modules ->
                pageModules model

            NotFound ->
                pageNotFound


pageHome : Model -> List (Html Msg)
pageHome model =
    [ h1 [] [ text "Home" ]
    , Grid.row []
        [ Grid.col []
            [ Card.config [ Card.outlinePrimary ]
                |> Card.headerH4 [] [ text "Getting started" ]
                |> Card.block []
                    [ Card.text [] [ text "Getting started is real easy. Just click the start button." ]
                    , Card.custom <|
                        Button.linkButton
                            [ Button.primary, Button.attrs [ href "#getting-started" ] ]
                            [ text "Start" ]
                    ]
                |> Card.view
            ]
        , Grid.col []
            [ Card.config [ Card.outlineDanger ]
                |> Card.headerH4 [] [ text "Modules" ]
                |> Card.block []
                    [ Card.text [] [ text "Check out the modules overview" ]
                    , Card.custom <|
                        Button.linkButton
                            [ Button.primary, Button.attrs [ href "#modules" ] ]
                            [ text "Module" ]
                    ]
                |> Card.view
            ]
        ]
    ]


pageSamples : Model -> List (Html Msg)
pageSamples model =
    [ h2 [] [ text "Samples" ]
    , ListGroup.custom
        (List.map
            (\sample ->
                ListGroup.anchor
                    [ ListGroup.attrs [ class "flex-column align-items-start" ] ]
                    (renderSample sample)
            )
            model.samples
        )
    ]


renderSample : Sample -> List (Html Msg)
renderSample sample =
    [ h5 [ class "mb-1" ] [ text sample.serialNumber ]
    , p [ class "mb-1" ]
        [ ul []
            [ li [] [ text ("cycle count: " ++ Basics.toString (sample.cycleCountValue)) ]
            , li [] [ text ("mode: " ++ Basics.toString (sample.mode)) ]
            , li [] [ text ("pre spray: " ++ Basics.toString (sample.preSprayValue)) ]
            , li [] [ text ("post spray: " ++ Basics.toString (sample.postSprayValue)) ]
            , li [] [ text ("program: " ++ Basics.toString (sample.programValue)) ]
            , li [] [ text ("spray value: " ++ Basics.toString (sample.sprayValue)) ]
            , li [] [ text ("air regulator: " ++ Basics.toString (sample.airRegulatorValueLoc)) ]
            ]
        ]
    ]


pageGettingStarted : Model -> List (Html Msg)
pageGettingStarted model =
    [ h2 [] [ text "Getting started" ]
    , Button.button
        [ Button.success
        , Button.large
        , Button.block
        , Button.attrs [ onClick <| ModalMsg Modal.visibleState ]
        ]
        [ text "Click me" ]
    ]


pageModules : Model -> List (Html Msg)
pageModules model =
    [ h1 [] [ text "Modules" ]
    , ListGroup.ul
        [ ListGroup.li [] [ text "Alert" ]
        , ListGroup.li [] [ text "Badge" ]
        , ListGroup.li [] [ text "Card" ]
        ]
    ]


pageNotFound : List (Html Msg)
pageNotFound =
    [ h1 [] [ text "Not found" ]
    , text "Sorry couldn't find that page"
    ]


modal : Model -> Html Msg
modal model =
    Modal.config ModalMsg
        |> Modal.small
        |> Modal.h4 [] [ text "Getting started ?" ]
        |> Modal.body []
            [ Grid.containerFluid []
                [ Grid.row []
                    [ Grid.col
                        [ Col.xs6 ]
                        [ text "Col 1" ]
                    , Grid.col
                        [ Col.xs6 ]
                        [ text "Col 2" ]
                    ]
                ]
            ]
        |> Modal.view model.modalState
