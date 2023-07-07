import streamlit as st
from page.n1_Network_Tools import page1_funcs  # Import the dictionary of subpage functions
from page.n2_Analysis_Tools import page2_funcs
#from page.n3_Mapping_Demo import page3
#from page.n4_DataFrame_Demo import page4

st.set_page_config(page_title="Shadow Suite", page_icon="🕸")

page_names_to_funcs = {
    "Main Page": "main_page",
    "Network Tools": page1_funcs,  # Use the dictionary as the value for 'Network Tool'
    "Analysis Tools": page2_funcs,
    #"Page 3": page3,
    #"Page 4": page4,
}

selected_page = st.sidebar.selectbox("Select a page", page_names_to_funcs.keys())

if selected_page == "Main Page":
    st.markdown(
    """
    # Welcome to Streamlit! 👋
    Streamlit is an open-source app framework built specifically for
    Machine Learning and Data Science projects.
    **👈 Select a demo from the sidebar** to see some examples
    of what Streamlit can do!
    ### Want to learn more?
    - Check out [streamlit.io](https://streamlit.io)
    - Jump into our [documentation](https://docs.streamlit.io)
    - Ask a question in our [community
        forums](https://discuss.streamlit.io)
    ### See more complex demos
    - Use a neural net to [analyze the Udacity Self-driving Car Image
        Dataset](https://github.com/streamlit/demo-self-driving)
    - Explore a [New York City rideshare dataset](https://github.com/streamlit/demo-uber-nyc-pickups)
    """
    )
elif isinstance(page_names_to_funcs[selected_page], dict):
    # If the selected page has subpages
    selected_subpage = st.sidebar.selectbox("Select a function", page_names_to_funcs[selected_page].keys())
    page_names_to_funcs[selected_page][selected_subpage]()  # Run the selected subpage function
else:
    # If the selected page doesn't have subpages, just run the function
    page_names_to_funcs[selected_page]()