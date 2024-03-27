class Message:
    def help(self) -> str:
        text = """\
            ?info : emailとpasswordを登録します。以後配信予定情報の取得が可能になります。
            ?videos : 登録しているYouTubeチャンネルの配信予定を確認できます。
            """
        return text